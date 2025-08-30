package controller

import (
	"blockchain/global"
	"blockchain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type buy struct {
	OrderID string `json:"orderid" binding:"required"`
}

func Buy(c *gin.Context) {
	var iBuy buy
	if err := c.ShouldBindJSON(&iBuy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "JSON解析失败，请核对填写信息",
		})
		return
	}

	var trans model.Transaction
	if err := global.DB.Model(&model.Transaction{}).Where("trans_id = ?", iBuy.OrderID).First(&trans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "查询出错",
		})
		return
	}

	var user model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", trans.Seller).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "查询出错",
		})
		return
	}

	var iUser model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", c.GetString("name")).First(&iUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "查询出错",
		})
		return
	}

	if iUser.Gocoin < trans.Price {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "Not Enough",
			"msg":   "资金不足",
		})
		return
	}

	if err := global.DB.Model(&user).Update("gocoin", user.Gocoin+trans.Price).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "更新失败",
		})
		return
	}

	if err := global.DB.Model(&iUser).Update("gocoin", iUser.Gocoin-trans.Price).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "更新失败",
		})
		return
	}

	if err := global.DB.Model(&trans).Update("Receiver", iUser.Name).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"error": "",
		"msg":   "交易成功",
	})

}
