package model

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Name    string `gorm:"uniqueIndex"`
	Data    string
	Profile string
	Rarity  string
}
