package dto

import "time"

type Transaction struct {
	CardID   string    `json:"card_id"`
	Seller   string    `json:"seller"`
	Receiver string    `json:"receiver"`
	TransID  string    `json:"trans_id"`
	CreateAt time.Time `json:"create_at"`

	Price int `json:"price"`
}
