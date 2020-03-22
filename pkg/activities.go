package pkg

import (
	"fmt"
	"net/http"

	"github.com/traaction/our-life-before-corona/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Activity struct {
	DB *gorm.DB
}

// Adds an activity
func (a Activity) Add(c *gin.Context) {
	var activity models.Activity
	err := c.BindJSON(&activity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(activity)

	if err := a.DB.Create(&activity).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, models.ReturnId{ID: activity.ID})
}

// Activity returns a list of activities from a given input string.
func (a Activity) List(c *gin.Context) {
	activities := make([]models.ReturnNameId, 0)
	query := "%" + c.Param("activity") + "%"

	g := a.DB.Table("activities").Select("ID, Name").Limit(10).Where("Name ILIKE ?", query).Find(&activities)
	if g.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, g.Error)
		return
	}

	c.JSON(http.StatusOK, activities)
}

// Activities returns a list of all activities we currently know.
func (a Activity) ListAll(c *gin.Context) {
	activities := make([]models.ReturnNameId, 0)

	g := a.DB.Table("activities").Select("ID, Name").Find(&activities)
	if g.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, g.Error)
		return
	}
	c.JSON(http.StatusOK, activities)
}
