package model

import "gorm.io/gorm"

type Card struct {
	gorm.Model

	Name    string
	Profile string
	HashID  string
	Owner   string
	Avatar  string

	Rarity string
	Attack int
	Blood  int
	Cost   int
	OnSale bool
}
