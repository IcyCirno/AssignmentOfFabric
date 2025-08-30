package controller

import (
	"blockchain/global"
	"blockchain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {

	var cards []model.Card

	if err := global.DB.Model(&model.Card{}).Where("owner = ?", c.GetString("name")).Find(&cards).Error; err != nil {
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
		"data":  cards,
	})
}
