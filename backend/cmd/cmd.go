package cmd

import (
	"blockchain/config"
	"blockchain/dto"
	"blockchain/fabric"
	"blockchain/global"
	"blockchain/model"
	"blockchain/router"
	"blockchain/utils"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Start() {

	config.InitConfig()

	fabric.InitFabric()

	global.Logger = config.InitLogger()

	db, err := config.InitDB()
	if err != nil {
		panic(fmt.Sprintf("Mysql Load Error: %v", err))
	}
	global.DB = db

	rdClient, err := config.InitRedis()
	if err != nil {
		panic(fmt.Sprintf("Redis Load Error: %v", err))
	}
	global.RedisClient = rdClient

	/*err = initDB()
	if err != nil {
		panic(fmt.Sprintf("Cards Load Error: %v", err))
	}*/

	router.InitRouter()

}

func initDB() error {

	stRoot, _ := os.Getwd()
	jsonPath := filepath.Join(stRoot, "card_db", "card.json")
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	var cards []dto.CardInit
	if err := json.Unmarshal(data, &cards); err != nil {
		return err
	}

	for _, card := range cards {
		c := model.Card{
			Name:    card.Name,
			Data:    utils.GenerateCardData(filepath.Join(stRoot, "card_db", card.Rarity, card.Location)),
			Profile: card.Profile,
			Rarity:  card.Rarity,
		}
		if err := global.DB.Save(&c).Error; err != nil {
			return err
		}
	}
	return nil
}
