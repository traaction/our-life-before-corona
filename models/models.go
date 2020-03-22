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

/*
Adding city: Khed Brahma
aa71ef3f-5d30-48a5-a622-3d26bdffa429
&{{0 0001-01-01 00:00:00 +0000 UTC 2020-03-22 10:54:03.928706764 +0100 CET m=+172.028280268 <nil>} Khed Brahma 1 aa71ef3f-5d30-48a5-a622-3d26bdffa429}
Adding city: Kheda
60e5ff47-47ad-4c65-89ee-036057a26228
&{{0 0001-01-01 00:00:00 +0000 UTC 2020-03-22 10:54:03.956203999 +0100 CET m=+172.055777487 <nil>} Kheda 1 60e5ff47-47ad-4c65-89ee-036057a26228}
Adding city: Khātra
83d8e64d-3f46-4cdc-86aa-9627af87b52d
&{{0 0001-01-01 00:00:00 +0000 UTC 2020-03-22 10:54:03.984582601 +0100 CET m=+172.084156090 <nil>} Khātra 1 83d8e64d-3f46-4cdc-86aa-9627af87b52d}
Adding city: Khatīma
a118d5be-9a7f-471c-a589-1bf4550ed62d
&{{0 0001-01-01 00:00:00 +0000 UTC 2020-03-22 10:54:04.008818807 +0100 CET m=+172.108392295 <nil>} Khatīma 1 a118d5be-9a7f-471c-a589-1bf4550ed62d}

Adding city: Khātegaon
34d9ba90-73a9-40b4-8287-3cd92ca71cc8
&{{0 0001-01-01 00:00:00 +0000 UTC 2020-03-22 10:54:04.035238124 +0100 CET m=+172.134811613 <nil>} Khātegaon 1 34d9ba90-73a9-40b4-8287-3cd92ca71cc8}
*/

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
