package test

import (
	"blockchain/dto"
	"blockchain/global"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCard(c *gin.Context) {
	var card dto.Card
	card.HashID = utils.GenerateOrderID()
	global.Logger.Info(card.HashID)
	if err := dto.PutCard(card); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "Fail", "")
		return
	}
	utils.Ok(c, "Ok", "")
}
