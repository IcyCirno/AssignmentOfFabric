package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var rarity = []string{"common", "rare", "epic", "legendary"}

func GenerateCardID(name string, profile string, owner string) string {
	info := fmt.Sprintf("%s-%s-%s-%d", name, profile, owner, time.Now().UnixNano())
	hash := sha256.Sum256([]byte(info))
	return hex.EncodeToString(hash[:])
}

func RandomInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

func RandomRarity() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rarity[r.Intn(4)]
}

func GenerateOrderID() string {
	return uuid.New().String()
}
