package dto

import "time"

type User struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`

	Rank   int `json:"rank"`
	Gocoin int `json:"gocoin"`

	Mine    bool      `json:"mine"`
	EndTime time.Time `json:"end_time"`

	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}
