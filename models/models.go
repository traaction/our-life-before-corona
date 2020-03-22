package models

import (
	"github.com/jinzhu/gorm"
)

type PlaceType int

const (
	Country PlaceType = iota
	City
	Other
)

type Place struct {
	gorm.Model
	Name string    `gorm:"type:varchar(100);NOT NULL" json:"name" binding:"required"`
	Type PlaceType `json:"-"`
}

type Activity struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);NOT NULL" json:"name" binding:"required"`
}

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type UserInfo struct {
	ID       string
	Name     string
	Location Location
}

type Sentence struct {
	gorm.Model
	Activity Activity `gorm:"foreignkey:activity_id;association_foreignkey:id"`
	AID      uint     `gorm:"column:activity_id"`
	Place    Place    `gorm:"foreignkey:place_id;association_foreignkey:id"`
	PID      uint     `gorm:"column:place_id"`
	UserInfo UserInfo
	Location Location
}

type PayloadSentence struct {
	UserID       string   `json:"userUuid" binding:"required"`
	UserName     string   `json:"userName" binding:"required"`
	ActivityID   uint     `json:"activityUuid" binding:"required"`
	PlaceID      uint     `json:"placeUuid" binding:"required"`
	UserLocation Location `json:"userLocation" binding:"required"`
}

type ReturnId struct {
	ID uint `json:"uuid"`
}

type ReturnNameId struct {
	ID   uint   `json:"uuid"`
	Name string `json:"name"`
}
