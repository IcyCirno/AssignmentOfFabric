package dto

import "time"

type Card struct {
	HashID   string    `json:"hash_id"`
	Owner    string    `json:"owner"`
	CreateAt time.Time `json:"create_at"`

	Name    string `json:"name"`
	Profile string `json:"profile"`

	Avatar string `json:"avatar"`
	Rarity string `json:"rarity"`

	Attack int `json:"attack"`
	Blood  int `json:"blood"`
	Cost   int `json:"cost"`

	OnSale    bool `json:"on_sale"`
	OnDefense bool `json:"on_defense"`
}
