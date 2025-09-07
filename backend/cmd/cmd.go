package cmd

import (
	"blockchain/config"
	"blockchain/fabric"
	"blockchain/global"
	"blockchain/model"
	"blockchain/router"
	"fmt"
	"time"
)

func Start() {

	config.InitConfig()

	fabric.InitFabric()

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

func Update() {

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		var users []model.User
		if err := global.DB.Find(&users).Error; err != nil {
			global.Logger.Error("查询数据库失败")
			continue
		}
		for _, user := range users {
			if user.Mine && time.Now().After(user.EndTime) {
				user.Mine = false
				if err := global.DB.Save(&user).Error; err != nil {
					global.Logger.Error("更新数据库失败")
				}
			}

		}

	}

}
