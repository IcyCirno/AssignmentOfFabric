package model

import "blockchain/dto"

// swagger:model CardAndTrans
type CardAndTrans struct {
	dto.Card
	dto.Transaction
}
