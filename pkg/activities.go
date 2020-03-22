package pkg

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/traaction/our-life-before-corona/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Activity struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

// Adds an activity
func (a Activity) Add(c *gin.Context) {
	var activity models.Activity
	err := c.BindJSON(&activity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	a.Logger.Info(activity)

	if err := a.DB.Create(&activity).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, activity.UUID)
}

// Activity returns a list of activities from a given input string.
func (a Activity) List(c *gin.Context) {
	activities := make([]models.Activity, 0)
	query := "%" + c.Param("activity") + "%"

	g := a.DB.Table("activities").Select("UUID, Name").Limit(10).Where("Name ILIKE ?", query).Find(&activities)
	if g.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, g.Error)
		return
	}

	c.JSON(http.StatusOK, activities)
}

// Activities returns a list of all activities we currently know.
func (a Activity) ListAll(c *gin.Context) {
	activities := make([]models.Activity, 0)

	g := a.DB.Table("activities").Select("UUID, Name").Find(&activities)
	if g.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, g.Error)
		return
	}
	c.JSON(http.StatusOK, activities)
}
