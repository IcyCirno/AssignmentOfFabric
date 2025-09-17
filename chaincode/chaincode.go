package main

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type User struct {
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`

	Rank   int `json:"rank"`
	Gocoin int `json:"gocoin"`

	Mine    bool      `json:"mine"`
	EndTime time.Time `json:"end_time"`

	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

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

type Transaction struct {
	CardID   string    `json:"card_id"`
	Seller   string    `json:"seller"`
	Receiver string    `json:"receiver"`
	TransID  string    `json:"trans_id"`
	CreateAt time.Time `json:"create_at"`

	Price int `json:"price"`
}

type SmartContract struct{ contractapi.Contract }

func (m *SmartContract) PutUser(c contractapi.TransactionContextInterface, userJSON string) error {
	var user User
	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		return err
	}
	return c.GetStub().PutState(user.Name, []byte(userJSON))
}

func (m *SmartContract) PutCard(c contractapi.TransactionContextInterface, cardJSON string) error {
	var card Card
	if err := json.Unmarshal([]byte(cardJSON), &card); err != nil {
		return err
	}
	return c.GetStub().PutState(card.HashID, []byte(cardJSON))
}

func (m *SmartContract) PutTransaction(c contractapi.TransactionContextInterface, transactionJSON string) error {
	var transaction Transaction
	if err := json.Unmarshal([]byte(transactionJSON), &transaction); err != nil {
		return err
	}
	return c.GetStub().PutState(transaction.TransID, []byte(transactionJSON))
}

func (m *SmartContract) GetUser(c contractapi.TransactionContextInterface, name string) (string, error) {
	data, err := c.GetStub().GetState(name)
	if err != nil || data == nil {
		return "", errors.New("Not Found")
	}
	return string(data), err
}

func (m *SmartContract) GetCard(c contractapi.TransactionContextInterface, hash_id string) (string, error) {
	data, err := c.GetStub().GetState(hash_id)
	if err != nil || data == nil {
		return "", errors.New("Not Found")
	}
	return string(data), err
}
func (m *SmartContract) GetTransaction(c contractapi.TransactionContextInterface, trans_id string) (string, error) {
	data, err := c.GetStub().GetState(trans_id)
	if err != nil || data == nil {
		return "", errors.New("Not Found")
	}
	return string(data), err
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		panic(err)
	}
	if err := chaincode.Start(); err != nil {
		panic(err)
	}
}
