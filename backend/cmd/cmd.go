package cmd

import (
	"blockchain/config"
	"blockchain/fabric"
	"blockchain/global"
	"blockchain/router"
	"fmt"
)

func Start() {

	config.InitConfig()

	fabric.InitFabric()

	global.Logger = config.InitLogger()

	rdClient, err := config.InitRedis()
	if err != nil {
		panic(fmt.Sprintf("Redis Load Error: %v", err))
	}
	global.RedisClient = rdClient

	router.InitRouter()

}
