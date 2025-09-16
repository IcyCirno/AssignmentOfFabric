package model

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);uniqueIndex"`
	Data    string
	Profile string
	Rarity  string
}
