package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type hashID struct {
	HashID string `json:"hashid" binding:"required"`
}

func Destroy(c *gin.Context) {
	var ihash hashID
	if err := c.ShouldBindJSON(&ihash); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	iCard, err := dto.GetCard(ihash.HashID)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值错误", nil)
		return
	}

	if iCard.OnSale {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌正在市场中", nil)
		return
	}

	iUser, err := dto.GetUser(iCard.Owner)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "无法找到", nil)
		return
	}

	iUser.Gocoin += viper.GetInt("nft.destroy")
	for i, id := range iUser.Cards {
		if id == iCard.HashID {
			iUser.Cards = append(iUser.Cards[:i], iUser.Cards[i+1:]...)
			break
		}
	}
	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "无法找到", nil)
		return
	}

	iCard.Destroy = true
	if err := dto.PutCard(iCard); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "删除失败", nil)
		return
	}

	utils.Ok(c, "摧毁成功", nil)

}
