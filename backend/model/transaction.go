package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	CardID   string
	Seller   string
	Receiver string
	TransID  string

	Price int
}
