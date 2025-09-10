package controller

import (
	"blockchain/dto"
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

	iUser, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", nil)
		return
	}

	if time.Now().Before(iUser.EndTime.Add(viper.GetDuration("nft.minetime") * time.Hour)) {
		utils.Fail(c, http.StatusBadRequest, "", "冷却中", nil)
		return
	}

	iCard, err := dto.GetCard(info.HashID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", nil)
		return
	}

	if iCard.OnSale {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌正在市场", nil)
		return
	}

	iUser.Gocoin += utils.RandomMine(iCard.Rarity)
	iUser.EndTime = time.Now().Add(time.Duration(viper.GetInt("nft.minetime")) * time.Hour)

	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	utils.Ok(c, "挖矿成功", nil)

}
