package dto

import (
	"blockchain/fabric"
	"encoding/json"
	"time"
)

type Card struct {
	HashID   string    `json:"hashid"`
	Owner    string    `json:"owner"`
	CreateAt time.Time `json:"createat"`

	Name    string `json:"name"`
	Profile string `json:"profile"`

	Avatar string `json:"avatar"`
	Rarity string `json:"rarity"`

	Attack int `json:"attack"`
	Blood  int `json:"blood"`
	Cost   int `json:"cost"`

	OnSale  bool   `json:"onsale"`
	TransID string `json:"transid"`

	OnDefense bool `json:"ondefense"`
	Destroy   bool `json:"destroy"`
}

func (m *Card) Free() bool {
	return !m.OnSale && !m.Destroy
}

func PutCard(card Card) error {
	data, err := json.Marshal(card)
	if err == nil {
		_, err = fabric.Contract.SubmitTransaction("PutCard", string(data))
	}
	return err
}

func GetCard(hash_id string) (Card, error) {
	data, err := fabric.Contract.EvaluateTransaction("GetCard", hash_id)
	var card Card
	if err == nil {
		err = json.Unmarshal(data, &card)
	}
	return card, err
}
