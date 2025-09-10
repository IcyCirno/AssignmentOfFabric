package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {

	var cards []dto.Card

	iUser, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", nil)
		return
	}

	for _, id := range iUser.Cards {
		card, err := dto.GetCard(id)
		if err != nil {
			utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", nil)
			return
		}
		cards = append(cards, card)
	}

	utils.Ok(c, "查询成功", cards)
}
