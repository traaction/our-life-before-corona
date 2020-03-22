package pkg

import (
	"fmt"
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
	c.JSON(http.StatusCreated, place.UUID)
}

// Place returns a list of places from a given input string.
func (p Place) List(c *gin.Context) {
	var places []models.ReturnNameId
	query := "%" + c.Param("place") + "%"
	orderStrPos := fmt.Sprintf("strpos(LOWER(Name), LOWER('%s')) ASC ", c.Param("place"))
	g := p.DB.Table("places").Select("UUID, Name").Where("Name ILIKE ?", query).Order(orderStrPos).Limit(10).Find(&places)

	if g.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, g.Error)
		return
	}

	c.JSON(http.StatusOK, places)
}

// Places returns a list of all places we currently know.
func (p Place) ListAll(c *gin.Context) {
	places := make([]models.ReturnNameId, 0)
	if err := p.DB.Table("places").Select("UUID, Name").Find(&places).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, places)
}
