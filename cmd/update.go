package cmd

import (
	"blockchain/global"
	"blockchain/model"
	"time"
)

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
