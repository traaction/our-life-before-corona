package pkg

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/traaction/our-life-before-corona/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Sentence struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

//Get all sentences fo user_id
func (s Sentence) Get(c *gin.Context) {
	var userInfo models.UserInfo
	s.Logger.Info(c.Param("user"))
	if err := s.DB.Where("uuid = ?", c.Param("user")).First(&userInfo).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}

	sentences := make([]models.Sentence, 0)
	if err := s.DB.Model(&userInfo).Preload("UserInfo").Preload("Activity").Preload("Place").Related(&sentences).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, sentences)
}

// Adds a Sentence
func (s Sentence) Add(c *gin.Context) {
	var payload models.PayloadSentence
	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	s.Logger.Info(payload)

	s.Logger.Info("Get Activity")
	var activity models.Activity
	if err := s.DB.Where("uuid = ?", payload.ActivityUUID).First(&activity).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}
	s.Logger.Info(activity.Name)

	s.Logger.Info("Get Place")
	var place models.Place
	if err := s.DB.Where("uuid = ?", payload.PlaceUUID).First(&place).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}
	s.Logger.Info(place.Name)

	s.Logger.Info("Get user info")
	var userInfo models.UserInfo
	if err := s.DB.Where("uuid = ?", payload.UserUUID).First(&userInfo).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			s.Logger.Info("Create user info")
			userInfo = models.UserInfo{UUID: payload.UserUUID, Lat: payload.UserLocation.Lat, Long: payload.UserLocation.Long}
			if err := s.DB.Create(&userInfo).Error; err != nil {
				s.Logger.Info(err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, err)
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

	}
	s.Logger.Info("Create Sentence")
	sentence := models.Sentence{ActivityID: activity.ID, PlaceID: place.ID, UserInfoID: userInfo.ID}

	if err := s.DB.Create(&sentence).Error; err != nil {
		s.Logger.Info(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, sentence.UUID)
}
