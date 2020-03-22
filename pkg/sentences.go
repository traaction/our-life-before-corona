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

	location := models.Location{Lat: payload.UserLocation.Lat, Long: payload.UserLocation.Long}

	userInfo := models.UserInfo{UUID: payload.UserUUID, Name: payload.UserName, Location: location}

	s.Logger.Info("Create Sentence")
	sentence := models.Sentence{Activity: activity, Place: place, UserInfo: userInfo}
	s.Logger.Info(sentence)
	if err := s.DB.Create(&sentence).Error; err != nil {
		s.Logger.Info(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, models.ReturnId{ID: sentence.UUID})
}
