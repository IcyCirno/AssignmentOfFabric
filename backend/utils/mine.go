package utils

import (
	"blockchain/dto"
	"math/rand"
)

var rarityBonus = map[string]float64{"common": 0.0, "rare": 0.05, "epic": 0.10, "legendary": 0.15}

func RandomMine(d string, a dto.Card, b dto.Card, c dto.Card) (bool, int) {
	totalStrength := a.Attack + a.Blood + b.Attack + b.Blood + c.Attack + c.Blood
	totalRarityBonus := rarityBonus[a.Rarity] + rarityBonus[b.Rarity] + rarityBonus[c.Rarity]
	avgStrength := float64(totalStrength) / 3.0

	var baseSuccessRate float64
	var rewardMultiplier float64

	switch d {
	case "simple":
		baseSuccessRate = 0.7
		rewardMultiplier = 1.0
	case "common":
		baseSuccessRate = 0.5
		rewardMultiplier = 2.0
	case "hard":
		baseSuccessRate = 0.3
		rewardMultiplier = 3.5
	}

	successRate := baseSuccessRate + avgStrength*0.002 + totalRarityBonus
	if successRate > 0.95 {
		successRate = 0.95
	}

	if rand.Float64() <= successRate {
		baseMoney := rand.Intn(10) + 5
		money := int(float64(baseMoney)*rewardMultiplier + avgStrength*0.5)

		return true, money
	} else {
		return false, 0
	}

}
