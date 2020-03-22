package models

import (
	"github.com/google/uuid"
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
	UUID uuid.UUID `gorm:"column:uuid;NOT NULL;UNIQUE"`
}

func (u *Place) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	err := scope.SetColumn("uuid", uuid)
	if err != nil {
		return err
	}
	return nil
}

type Activity struct {
	gorm.Model
	Name string    `gorm:"type:varchar(100);NOT NULL" json:"name" binding:"required"`
	UUID uuid.UUID `gorm:"column:uuid;NOT NULL;UNIQUE"`
}

func (u *Activity) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	err := scope.SetColumn("uuid", uuid)
	if err != nil {
		return err
	}
	return nil
}

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type UserInfo struct {
	UUID     uuid.UUID
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
	UUID     uuid.UUID `gorm:"column:uuid;NOT NULL;UNIQUE"`
}

func (u *Sentence) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	err := scope.SetColumn("uuid", uuid)
	if err != nil {
		return err
	}
	return nil
}

type PayloadSentence struct {
	UserUUID     uuid.UUID `json:"userUuid" binding:"required"`
	UserName     string    `json:"userName" binding:"required"`
	ActivityUUID uuid.UUID `json:"activityUuid" binding:"required"`
	PlaceUUID    uuid.UUID `json:"placeUuid" binding:"required"`
	UserLocation Location  `json:"userLocation" binding:"required"`
}

type ReturnId struct {
	ID uuid.UUID `json:"uuid"`
}

type ReturnNameId struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
