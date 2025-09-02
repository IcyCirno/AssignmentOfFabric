package model

import "gorm.io/gorm"

type Card struct {
	gorm.Model

	Name    string `json:"name"`
	Profile string `json:"profile"`
	HashID  string `json:"hashid"`
	Owner   string `json:"owner"`
	Avatar  string `json:"avatar"`

	Rarity string `json:"rarity"`
	Attack int    `json:"attack"`
	Blood  int    `json:"blood"`
	Cost   int    `json:"cost"`
	OnSale bool   `json:"onsale"`
}
