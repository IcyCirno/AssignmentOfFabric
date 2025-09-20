package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type cardCreate struct {
	Name    string `json:"name" binding:"required"`
	Profile string `json:"profile"`
	Invest  int    `json:"invest" binding:"required"`
}

func Mint(c *gin.Context) {

	var info cardCreate
	if err := c.ShouldBindJSON(&info); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", "")
		return
	}

	owner := c.GetString("name")
	iUser, err := dto.GetUser(owner)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "用户无法从区块链中获得", "")
		return
	}

	if iUser.Gocoin < viper.GetInt("nft.mintcost") {
		utils.Fail(c, http.StatusBadRequest, "", "资金不足", "")
		return
	}

	iCard, err := utils.CreateCard(info.Name, c.GetString("name"), info.Invest)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器错误", "")
		return
	}

	iUser.Gocoin -= info.Invest
	iUser.Cards = append(iUser.Cards, iCard.HashID)

	if err := dto.PutCard(iCard); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "添加卡牌出错", "")
		return
	}

	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "修改用户出错", "")
		return
	}

	utils.Ok(c, "铸造成功", iCard)
}
