package dto

type CardInit struct {
	Name     string `json:"name" binding:"required"`
	Profile  string `json:"profile" binding:"required"`
	Rarity   string `json:"rarity" binding:"required"`
	Location string `json:"location" binding:"required"`
}
