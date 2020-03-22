package pkg

import (
	"fmt"
	"net/http"

	"github.com/traaction/our-life-before-corona/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Sentence struct {
	DB *gorm.DB
}

func (s Sentence) Get(c *gin.Context) {
	var sentence models.Sentence
	fmt.Println(c.Param("sentence"))
	if err := s.DB.Preload("UserInfo").Preload("Activity").Preload("Place").Where("uuid = ?", c.Param("sentence")).First(&sentence).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}
	fmt.Println(sentence)

	c.JSON(http.StatusOK, sentence)
}

// Adds a Sentence
func (s Sentence) Add(c *gin.Context) {
	var payload models.PayloadSentence
	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	fmt.Println(payload)

	fmt.Println("Get Activity")
	var activity models.Activity
	if err := s.DB.Where("uuid = ?", payload.ActivityUUID).First(&activity).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}
	fmt.Println(activity.Name)

	fmt.Println("Get Place")
	var place models.Place
	if err := s.DB.Where("uuid = ?", payload.PlaceUUID).First(&place).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}
	fmt.Println(place.Name)

	fmt.Println("Get user info")
	var userInfo models.UserInfo
	if err := s.DB.Where("uuid = ?", payload.UserUUID).First(&userInfo).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Create user info")
			userInfo = models.UserInfo{UUID: payload.UserUUID, Lat: payload.UserLocation.Lat, Long: payload.UserLocation.Long}
			if err := s.DB.Create(&userInfo).Error; err != nil {
				fmt.Println(err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, err)
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

	}

	fmt.Println("Create Sentence")
	sentence := models.Sentence{ActivityID: activity.ID, PlaceID: place.ID, UserInfoID: userInfo.ID}
	if err := s.DB.Create(&sentence).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, sentence.UUID)
}
