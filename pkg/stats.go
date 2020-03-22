package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/traaction/our-life-before-corona/models"

	"github.com/jinzhu/gorm"
)

type Stats struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

// GEt returns a number stats from a given input sentence UUID.
func (s Stats) Get(c *gin.Context) {
	var payload models.PayloadStats

	err := c.BindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var sentence models.Sentence
	g := s.DB.Table("sentences").Preload("UserInfo").Preload("Activity").Preload("Place").Where("uuid = ?", payload.SentenceUUID).First(&sentence)
	if g.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, g.Error)
		return
	}

	stats := &models.Stats{}
	err = s.calculateActivityStats(sentence.Activity, stats)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	err = s.calculatePlaceStats(sentence.Place, stats)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (s Stats) calculateActivityStats(activitiy models.Activity, stats *models.Stats) error {
	totalCount := 0
	matchCount := 0

	g := s.DB.Table("sentences").Where("activity_ID = ?", activitiy.ID).Count(&matchCount)
	if g.Error != nil {
		return g.Error
	}
	stats.ActivityStats.MatchCount = matchCount

	g = s.DB.Table("sentences").Count(&totalCount)
	if g.Error != nil {
		return g.Error
	}
	stats.ActivityStats.TotalCount = totalCount

	return nil
}

func (s Stats) calculatePlaceStats(activitiy models.Place, stats *models.Stats) error {
	totalCount := 0
	matchCount := 0

	g := s.DB.Table("sentences").Where("place_ID = ?", activitiy.ID).Count(&matchCount)
	if g.Error != nil {
		return g.Error
	}
	stats.PlaceStats.MatchCount = matchCount

	g = s.DB.Table("sentences").Count(&totalCount)
	if g.Error != nil {
		return g.Error
	}
	stats.PlaceStats.TotalCount = totalCount

	return nil
}
