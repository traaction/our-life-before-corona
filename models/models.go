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
	gorm.Model `json:"-"`
	Name       string     `gorm:"type:varchar(100);NOT NULL" json:"name"`
	Type       PlaceType  `gorm:"type:int" json:"-"`
	UUID       uuid.UUID  `gorm:"column:uuid;NOT NULL;UNIQUE" json:"uuid"`
	Sentences  []Sentence `json:"-"`
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
	gorm.Model `json:"-"`
	Name       string     `gorm:"type:varchar(100);NOT NULL" json:"name"`
	UUID       uuid.UUID  `gorm:"column:uuid;NOT NULL;UNIQUE" json:"uuid"`
	Sentences  []Sentence `json:"-"`
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
	gorm.Model `json:"-"`
	UUID       uuid.UUID  `gorm:"column:uuid;NOT NULL" json:"userUUID"`
	Lat        float64    `json:"lat"`
	Long       float64    `json:"long"`
	Sentences  []Sentence `json:"-"`
}

type Sentence struct {
	gorm.Model `json:"-"`
	ActivityID uint      `json:"-"`
	Activity   Activity  `json:"activity"`
	PlaceID    uint      `json:"-"`
	Place      Place     `json:"place"`
	UserInfoID uint      `json:"-"`
	UserInfo   UserInfo  `json:"userInfo"`
	UUID       uuid.UUID `gorm:"column:uuid;NOT NULL;UNIQUE" json:"uuid"`
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
	ActivityUUID uuid.UUID `json:"activityUuid" binding:"required"`
	PlaceUUID    uuid.UUID `json:"placeUuid" binding:"required"`
	UserLocation Location  `json:"userLocation" binding:"required"`
}

type Stats struct {
	ActivityStats StatsActivity ` json:"activity"`
	PlaceStats    StatsPlace    ` json:"place"`
	SentenceStats StatsSentence ` json:"sentence"`
}

type StatsActivity struct {
	TotalDistinctCount int `json:"totalDistinctCount"`
	MatchCount         int `json:"matchCount"`
}

type StatsPlace struct {
	TotalDistinctCount int `json:"totalDistinctCount"`
	MatchCount         int `json:"matchCount"`
}

type StatsSentence struct {
	TotalDistinctCount int `json:"totalDistinctCount"`
}

type PayloadStats struct {
	SentenceUUID uuid.UUID `json:"sentenceUuid" binding:"required"`
}
