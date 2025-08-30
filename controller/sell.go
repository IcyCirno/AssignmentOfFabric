package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sell struct {
	HashID string `json:"hashid" binding:"required"`
	Cost   int    `json:"cost" binding:"required"`
}

func Sell(c *gin.Context) {
	var isell sell
	if err := c.ShouldBindJSON(&isell); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "JSON解析失败，请核对填写信息",
		})
		return
	}

	var iCard model.Card
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", isell.HashID).First(&iCard).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "哈希值出错",
		})
		return
	}

	trans := model.Transaction{
		CardID:  iCard.HashID,
		Seller:  iCard.Owner,
		TransID: utils.GenerateOrderID(),
		Price:   isell.Cost,
	}

	if err := global.DB.Model(&iCard).Update("on_sale", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "无法更新状态",
		})
		return
	}

	if err := global.DB.Save(&trans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "无法创建订单",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"error": "",
		"msg":   "创建交易成功",
		"data":  trans,
	})

}
