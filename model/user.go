package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string
	Email    string
	Password string

	Rank   int
	Gocoin int

	Mine    bool
	EndTime time.Time

	A string
	B string
	C string
}
