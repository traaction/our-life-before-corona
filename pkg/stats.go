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

	err = s.calculateSentenceStats(sentence, stats)
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

	g = s.DB.Table("activities").Count(&totalCount)
	if g.Error != nil {
		return g.Error
	}
	stats.ActivityStats.TotalDistinctCount = totalCount

	return nil
}

func (s Stats) calculatePlaceStats(place models.Place, stats *models.Stats) error {
	totalCount := 0
	matchCount := 0

	g := s.DB.Table("sentences").Where("place_ID = ?", place.ID).Count(&matchCount)
	if g.Error != nil {
		return g.Error
	}
	stats.PlaceStats.MatchCount = matchCount

	g = s.DB.Table("places").Count(&totalCount)
	if g.Error != nil {
		return g.Error
	}
	stats.PlaceStats.TotalDistinctCount = totalCount

	return nil
}

func (s Stats) calculateSentenceStats(sentence models.Sentence, stats *models.Stats) error {
	totalCount := 0
	matchCount := 0

	g := s.DB.Table("sentences").Count(&totalCount)
	if g.Error != nil {
		return g.Error
	}
	stats.SentenceStats.TotalDistinctCount = totalCount

	g = s.DB.Table("sentences").Where("place_ID = ? AND activity_ID = ?", sentence.Place.ID, sentence.Activity.ID).Count(&matchCount)
	if g.Error != nil {
		return g.Error
	}
	stats.SentenceStats.MatchCount = matchCount

	return nil
}
