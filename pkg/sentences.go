package pkg

import (
	"fmt"
	"gitlab/wirvsvirus/our-life-before-corona/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Sentence struct {
	DB *gorm.DB
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
	if err := s.DB.First(&activity, payload.ActivityID).Error; err != nil {
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
	if err := s.DB.First(&place, payload.PlaceID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}
	fmt.Println(place.Name)

	location := models.Location{Lat: payload.UserLocation.Lat, Long: payload.UserLocation.Long}

	userInfo := models.UserInfo{ID: payload.UserID, Name: payload.UserName, Location: location}

	fmt.Println("Create Sentence")
	sentence := models.Sentence{Activity: activity, Place: place, UserInfo: userInfo}
	fmt.Println(sentence)
	if err := s.DB.Create(&sentence).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, models.ReturnId{ID: sentence.ID})
}
