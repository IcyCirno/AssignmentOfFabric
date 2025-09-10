package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// cardCreate 铸造卡牌请求参数
// swagger:model cardCreate
type cardCreate struct {
	// 卡牌名称
	Name string `json:"name" binding:"required"`
	// 卡牌描述
	Profile string `json:"profile" binding:"required"`
	// 投资等级，用于随机稀有度
	Invest int `json:"invest" binding:"required"`
}

// Mint godoc
// @Summary 铸造卡牌
// @Description 用户使用平台货币铸造新的卡牌，生成随机属性和稀有度
// @Tags NFT
// @Accept json
// @Produce json
// @Param data body cardCreate true "卡牌铸造信息"
// @Success 200 {object} utils.APIResponse[dto.Card] "铸造成功，返回卡牌信息"
// @Failure 400 {object} utils.APIResponse[string] "请求参数错误或资金不足"
// @Failure 500 {object} utils.APIResponse[string] "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/auth/card/mint [post]
func Mint(c *gin.Context) {

	var info cardCreate
	if err := c.ShouldBindJSON(&info); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", "")
		return
	}

	owner := c.GetString("name")
	iUser, err := dto.GetUser(owner)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错", "")
		return
	}

	if iUser.Gocoin < viper.GetInt("nft.mintcost") {
		utils.Fail(c, http.StatusBadRequest, "", "资金不足", "")
		return
	}

	iCard := dto.Card{
		Name:    info.Name,
		Profile: info.Profile,
		HashID:  utils.GenerateCardID(info.Name, info.Profile, owner),
		Owner:   owner,

		Attack: utils.RandomAttack(),
		Blood:  utils.RandomBlood(),
		Cost:   utils.RandomCost(),
		Rarity: utils.RandomRarity(info.Invest),

		OnSale:    false,
		OnDefense: false,
		Destroy:   false,
	}
	iCard.Avatar = utils.RandomAvatar(iCard.Rarity)

	iUser.Gocoin -= viper.GetInt("nft.mintcost")
	iUser.Cards = append(iUser.Cards, iCard.HashID)

	if err := dto.PutCard(iCard); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错", "")
		return
	}

	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错", "")
		return
	}

	utils.Ok(c, "铸造成功", iCard)
}
