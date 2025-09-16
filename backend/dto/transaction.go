package dto

import (
	"blockchain/fabric"
	"encoding/json"
	"time"
)

// swagger:model Transaction
type Transaction struct {
	CardID   string    `json:"card_id"`
	Seller   string    `json:"seller"`
	Receiver string    `json:"receiver"`
	TransID  string    `json:"trans_id"`
	CreateAt time.Time `json:"create_at"`

	Price  int    `json:"price"`
	Status string `json:"status"`
}

func PutTransaction(transaction Transaction) error {
	data, err := json.Marshal(transaction)
	if err == nil {
		_, err = fabric.Contract.SubmitTransaction("PutTransaction", string(data))
	}
	return err
}

func GetTransaction(trans_id string) (Transaction, error) {
	data, err := fabric.Contract.EvaluateTransaction("GetTransaction", trans_id)
	var transaction Transaction
	if err == nil {
		err = json.Unmarshal(data, &transaction)
	}
	return transaction, err
}
