package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CancelRequest 取消交易请求参数
// swagger:model CancelRequest
type cancel struct {
	// 交易订单ID
	OrderID string `json:"orderid" binding:"required"`
}

// Cancel godoc
// @Summary 取消交易
// @Description 用户取消自己发起的卡牌交易订单
// @Tags NFT
// @Accept json
// @Produce json
// @Param data body cancel true "交易订单ID"
// @Success 200 {object} utils.APIResponse[string] "取消成功"
// @Failure 400 {object} utils.APIResponse[string] "请求参数错误"
// @Failure 500 {object} utils.APIResponse[string] "服务器内部错误或更新失败"
// @Security ApiKeyAuth
// @Router /api/auth/card/cancel [post]
func Cancel(c *gin.Context) {
	var icancel cancel
	if err := c.ShouldBind(&icancel); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", "")
		return
	}

	trans, err := dto.GetTransaction(icancel.OrderID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", "")
		return
	}

	iCard, err := dto.GetCard(trans.CardID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", "")
		return
	}

	iUser, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", "")
		return
	}

	iCard.OnSale = false
	if err := dto.PutCard(iCard); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", "")
		return
	}

	trans.Status = "Canceled"
	if err := dto.PutTransaction(trans); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "删除失败", "")
		return
	}

	for i, id := range iUser.Trans {
		if id == trans.TransID {
			iUser.Trans = append(iUser.Trans[:i], iUser.Trans[:i+1]...)
			break
		}
	}
	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", "")
		return
	}

	utils.Ok(c, "已取消", "")
}
