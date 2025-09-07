package controller

import (
	"blockchain/global"
	"blockchain/model"
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

	var trans model.Transaction
	if err := global.DB.Model(&model.Transaction{}).Where("trans_id = ?", icancel.OrderID).First(&trans).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", nil)
		return
	}

	var iCard model.Card
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", trans.CardID).First(&iCard).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询失败", nil)
		return
	}

	if err := global.DB.Model(&iCard).Update("on_sale", false).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	if err := global.DB.Delete(&trans).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "删除失败", nil)
		return
	}

	utils.Ok(c, "已取消", nil)

}
