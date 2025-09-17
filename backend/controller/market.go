package controller

import (
	"blockchain/dto"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Market godoc
// @Summary 查询市场交易
// @Description 查询当前登录用户参与的所有卡牌交易订单
// @Tags NFT
// @Accept json
// @Produce json
// @Success 200 {object} utils.APIResponse[[]model.CardAndTrans] "查询成功，返回交易列表"
// @Failure 500 {object} utils.APIResponse[string] "服务器内部错误或查询失败"
// @Security ApiKeyAuth
// @Router /api/auth/market/query [post]
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
