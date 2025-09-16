package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Query godoc
// @Summary 查询用户卡牌
// @Description 查询当前登录用户所拥有的所有卡牌信息
// @Tags NFT
// @Accept json
// @Produce json
// @Success 200 {object} utils.APIResponse[[]dto.Card] "查询成功，返回用户卡牌列表"
// @Failure 500 {object} utils.APIResponse[string] "服务器内部错误或查询失败"
// @Security ApiKeyAuth
// @Router /api/auth/card/query [post]
func Query(c *gin.Context) {

	var cards []dto.Card

	iUser, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", "")
		return
	}

	for _, id := range iUser.Cards {
		card, err := dto.GetCard(id)
		if err != nil {
			utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", "")
			return
		}
		cards = append(cards, card)
	}

	utils.Ok(c, "查询成功", cards)
}
