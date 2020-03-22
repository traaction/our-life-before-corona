package pkg

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/traaction/our-life-before-corona/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Place struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

// AddPlace adds a place to our datastore.
func (p Place) Add(c *gin.Context) {
	var place models.Place
	err := c.BindJSON(&place)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	place.Type = models.Other
	p.Logger.Info(place)

	if err := p.DB.Create(&place).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, place)
}

// Place returns a list of places from a given input string.
func (p Place) List(c *gin.Context) {
	places := make([]models.Place, 0)
	query := "%" + c.Param("place") + "%"
	if err := p.DB.Table("places").Select("UUID, Name").Limit(10).Where("Name ILIKE ?", query).Find(&places).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, places)
}

// Places returns a list of all places we currently know.
func (p Place) ListAll(c *gin.Context) {
	places := make([]models.Place, 0)
	if err := p.DB.Table("places").Select("UUID, Name").Find(&places).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, places)
}
