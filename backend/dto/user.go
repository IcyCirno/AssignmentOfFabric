package dto

import (
	"blockchain/fabric"
	"encoding/json"
	"time"
)

type User struct {
	Name     string    `json:"name"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`

	Rank   int `json:"rank"`
	Gocoin int `json:"gocoin"`

	EndTime time.Time `json:"end_time"`

	Cards []string
	Trans []string
}

func PutUser(user User) error {
	data, err := json.Marshal(user)
	if err == nil {
		_, err = fabric.Contract.SubmitTransaction("PutUser", string(data))
	}
	return err
}

func GetUser(name string) (User, error) {
	data, err := fabric.Contract.EvaluateTransaction("GetUser", name)
	var user User
	if err == nil {
		err = json.Unmarshal(data, &user)
	}
	return user, err
}
