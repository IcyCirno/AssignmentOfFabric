package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type mine struct {
	HashID string
}

func Mine(c *gin.Context) {
	var info mine

	if err := c.ShouldBind(&info); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	var iUser model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", c.GetString("name")).First(&iUser).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	if iUser.Mine {
		utils.Fail(c, http.StatusBadRequest, "", "冷却中", nil)
		return
	}

	var iCard model.Card
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", info.HashID).First(&iCard).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	if iCard.OnSale {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌正在市场", nil)
		return
	}

	iUser.Mine = true
	iUser.Gocoin += utils.RandomMine(iCard.Rarity)
	iUser.EndTime = time.Now().Add(time.Duration(viper.GetInt("nft.minetime")) * time.Hour)

	if err := global.DB.Save(&iUser).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	utils.Ok(c, "挖矿成功", nil)

}
