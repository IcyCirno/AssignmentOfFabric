package controller

import (
	"blockchain/dto"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SellRequest 上架卡牌请求参数
// swagger:model SellRequest
type sell struct {
	// 卡牌唯一ID
	HashID string `json:"hashid" binding:"required"`
	// 交易价格
	Cost int `json:"cost" binding:"required"`
}

// Sell godoc
// @Summary 上架卡牌
// @Description 用户将自己的卡牌上架市场进行交易
// @Tags NFT
// @Accept json
// @Produce json
// @Param data body sell true "卡牌上架信息"
// @Success 200 {object} utils.APIResponse[model.CardAndTrans] "创建交易成功，返回交易信息"
// @Failure 400 {object} utils.APIResponse[string] "请求参数错误或卡牌在市场中"
// @Failure 500 {object} utils.APIResponse[string] "服务器内部错误或创建交易失败"
// @Security ApiKeyAuth
// @Router /api/auth/card/sell [post]
func Sell(c *gin.Context) {
	var isell sell
	if err := c.ShouldBindJSON(&isell); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", "")
		return
	}

	if isell.Cost < 0 {
		utils.Fail(c, http.StatusBadRequest, "", "交易金额不合法", "")
		return
	}

	iCard, err := dto.GetCard(isell.HashID)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值出错", "")
		return
	}

	if iCard.OnSale {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌正在市场中", "")
		return
	}

	if iCard.Destroy {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌已被摧毁", "")
		return
	}

	iUser, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", "")
	}

	root, err := dto.GetUser("root")
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", "")
	}

	trans := dto.Transaction{
		CardID:  iCard.HashID,
		Seller:  iCard.Owner,
		TransID: utils.GenerateOrderID(),
		Price:   isell.Cost,
		Status:  "For sale",
	}

	iCard.OnSale = true
	iCard.TransID = trans.TransID
	if err := dto.PutCard(iCard); err != nil {
		utils.Fail(c, http.StatusInternalServerError, "无法更新", "无法更新", "")
		return
	}

	if err := dto.PutTransaction(trans); err != nil {
		utils.Fail(c, http.StatusInternalServerError, "无法创建订单", "无法创建订单", "")
		return
	}

	iUser.Trans = append(iUser.Trans, trans.TransID)
	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, "无法更新", "无法更新", "")
		return
	}

	root.Trans = append(root.Trans, trans.TransID)
	if err := dto.PutUser(root); err != nil {
		utils.Fail(c, http.StatusInternalServerError, "无法更新", "无法更新", "")
		return
	}

	utils.Ok(c, "创建交易成功", model.CardAndTrans{
		Card:        iCard,
		Transaction: trans,
	})
}
