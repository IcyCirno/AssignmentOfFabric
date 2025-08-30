package controller

import (
	"blockchain/global"
	"blockchain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Market(c *gin.Context) {
	var transactions []model.Transaction
	if err := global.DB.Model(&model.Transaction{}).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "查询出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"error": "",
		"msg":   "查询成功",
		"data":  transactions,
	})

}
