package utils

import (
	"blockchain/dto"
	"blockchain/global"
	"blockchain/model"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type rarity struct {
	Name   string
	Weight float64
}

const factor = 0.01

var (
	mine = map[string]int{"common": 1, "rare": 2, "epic": 3, "legendary": 4}
	list = []rarity{{"common", 25}, {"rare", 25}, {"epic", 25}, {"legendary", 25}}
)

func GenerateCardID(name string, profile string, owner string) string {
	info := fmt.Sprintf("%s-%s-%s-%d", name, profile, owner, time.Now().UnixNano())
	hash := sha256.Sum256([]byte(info))
	return hex.EncodeToString(hash[:])
}

func GenerateOrderID() string {
	return uuid.New().String()
}

func RandomAttack() int {
	return rand.IntN(100)
}

func RandomBlood() int {
	return rand.IntN(100)
}

func RandomCost() int {
	return rand.IntN(10)
}

func RandomRarity(invest int) string {
	var total float64
	w := make([]float64, len(list))
	for i, r := range list {
		b := 1 + factor*float64(invest)*float64(i)
		w[i] = r.Weight * b
		total += w[i]
	}

	roll := rand.Float64() * total
	t := 0.0
	for i, c := range w {
		t += c
		if t > roll {
			return list[i].Name
		}
	}
	return list[len(list)-1].Name
}

func RandomAvatar(r string) (model.Card, error) {
	var card model.Card
	err := global.DB.Where("rarity = ?", r).Order("RAND()").First(&card).Error
	return card, err
}

func GenerateCardData(imgPath string) string {
	data, _ := os.ReadFile(imgPath)
	encoded := base64.StdEncoding.EncodeToString(data)

	ext := strings.ToLower(filepath.Ext(imgPath))
	var mime string
	switch ext {
	case ".jpg", ".jpeg":
		mime = "image/jpeg"
	case ".png":
		mime = "image/png"
	case ".gif":
		mime = "image/gif"
	default:
		mime = "application/octet-stream"
	}

	return fmt.Sprintf("data:%s;base64,%s", mime, encoded)
}

func CreateCard(name string, owner string, invest int) (dto.Card, error) {
	iCard := dto.Card{
		Name:   name,
		HashID: GenerateCardID(name, "nft_card", owner),
		Owner:  owner,

		Attack: RandomAttack(),
		Blood:  RandomBlood(),
		Cost:   RandomCost(),
		Rarity: RandomRarity(invest),

		OnSale:    false,
		OnDefense: false,
		Destroy:   false,
	}
	temp, err := RandomAvatar(iCard.Rarity)
	iCard.Profile = temp.Profile
	iCard.Avatar = temp.Data

	return iCard, err
}
