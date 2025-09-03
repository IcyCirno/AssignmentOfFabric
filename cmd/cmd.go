package cmd

import (
	"blockchain/config"
	"blockchain/global"
	"blockchain/router"
	"fmt"
)

func Start() {

	config.InitConfig()

	global.Logger = config.InitLogger()

	db, err := config.InitDB()
	if err != nil {
		panic(fmt.Sprintf("DB Load Error: %v", err))
	}
	global.DB = db

	rdClient, err := config.InitRedis()
	if err != nil {
		panic(fmt.Sprintf("Redis Load Error: %v", err))
	}
	global.RedisClient = rdClient

	router.InitRouter()

}
