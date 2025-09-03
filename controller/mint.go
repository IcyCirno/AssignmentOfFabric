package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type cardCreate struct {
	Name    string `json:"name" binding:"required"`
	Profile string `json:"profile" binding:"required"`
	Invest  int    `json:"invest" bindinng:"required"`
}

func Mint(c *gin.Context) {

	var info cardCreate
	if err := c.ShouldBindJSON(&info); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}
	owner := c.GetString("name")
	var iUser model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", owner).First(&iUser).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错", nil)
		return
	}

	if iUser.Gocoin < viper.GetInt("nft.mintcost") {
		utils.Fail(c, http.StatusBadRequest, "", "资金不足", nil)
		return
	}

	iCard := model.Card{
		Name:    info.Name,
		Profile: info.Profile,
		HashID:  utils.GenerateCardID(info.Name, info.Profile, owner),
		Owner:   owner,

		Attack: utils.RandomAttack(),
		Blood:  utils.RandomBlood(),
		Cost:   utils.RandomCost(),
		Rarity: utils.RandomRarity(info.Invest),

		OnSale: false,
	}
	iCard.Avatar = utils.RandomAvatar(iCard.Rarity)

	//先上链

	if err := global.DB.Save(&iCard).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错", nil)
		return
	}

	if err := global.DB.Model(&iUser).Update("gocoin", iUser.Gocoin-viper.GetInt("nft.mintcost")).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错", nil)
		return
	}

	utils.Ok(c, "铸造成功", iCard)
}
