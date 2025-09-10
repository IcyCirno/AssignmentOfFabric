package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cancel struct {
	OrderID string `json:"orderid" binding:"required"`
}

func Cancel(c *gin.Context) {
	var icancel cancel
	if err := c.ShouldBind(&icancel); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	trans, err := dto.GetTransaction(icancel.OrderID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", nil)
		return
	}

	iCard, err := dto.GetCard(trans.CardID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", nil)
		return
	}

	iUser, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", nil)
		return
	}

	iCard.OnSale = false
	if err := dto.PutCard(iCard); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	trans.Status = "Canceled"
	if err := dto.PutTransaction(trans); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "删除失败", nil)
		return
	}

	for i, id := range iUser.Trans {
		if id == trans.TransID {
			iUser.Trans = append(iUser.Trans[:i], iUser.Trans[:i+1]...)
			break
		}
	}
	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	utils.Ok(c, "已取消", nil)

}
