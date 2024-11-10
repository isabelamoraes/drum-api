package model

import "gorm.io/gorm"

type Drum struct {
	gorm.Model
	Name          string
	Mark          string
	Configuration string
	IsEletronic   bool     `gorm:"default:false"`
	Reviews       []Review `gorm:"foreignKey:DrumID"`
}
