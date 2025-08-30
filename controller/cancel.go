package controller

import (
	"blockchain/global"
	"blockchain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cancel struct {
	OrderID string `json:"orderid" binding:"required"`
}

func Cancel(c *gin.Context) {
	var icancel cancel
	if err := c.ShouldBind(&icancel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "JSON解析失败，请核对填写信息",
		})
		return
	}

	var trans model.Transaction
	if err := global.DB.Model(&model.Transaction{}).Where("trans_id = ?", icancel.OrderID).First(&trans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "查询失败",
		})
		return
	}

	var iCard model.Card
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", trans.CardID).First(&iCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "查询失败",
		})
		return
	}

	if err := global.DB.Model(&iCard).Update("on_sale", false).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "更新失败",
		})
		return
	}

	if err := global.DB.Delete(&trans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "删除",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"error": "",
		"msg":   "已取消",
	})

}
