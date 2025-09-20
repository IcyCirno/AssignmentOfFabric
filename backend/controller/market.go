package controller

import (
	"blockchain/dto"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Market(c *gin.Context) {
	var transactions []model.CardAndTrans

	iUser, err := dto.GetUser("root")
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", "")
		return
	}

	for _, id := range iUser.Trans {
		trans, err := dto.GetTransaction(id)
		if err != nil {
			utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", "")
			return
		}
		card, err := dto.GetCard(trans.CardID)
		if err != nil {
			utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", "")
			return
		}
		transactions = append(transactions, model.CardAndTrans{
			Card:        card,
			Transaction: trans,
		})
	}

	utils.Ok(c, "查询成功", transactions)
}
