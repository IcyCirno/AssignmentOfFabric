package dto

import (
	"blockchain/fabric"
	"encoding/json"
	"time"
)

// swagger:model Card
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
	Destroy   bool `json:"destroy"`
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
