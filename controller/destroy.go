package controller

import (
	"blockchain/global"
	"blockchain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type hashID struct {
	HashID string `json:"hashid" binding:"required"`
}

func Destroy(c *gin.Context) {
	var ihash hashID
	if err := c.ShouldBindJSON(&ihash); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "JSON解析失败，请核对填写信息",
		})
		return
	}

	var iCard model.Card
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", ihash.HashID).First(&iCard).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "哈希值错误",
		})
		return
	}

	var iUser model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", iCard.Owner).First(&iUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "无法找到",
		})
		return
	}

	if err := global.DB.Model(&iUser).Update("gocoin", iUser.Gocoin+4).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "无法找到",
		})
		return
	}

	if err := global.DB.Delete(&iCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "无法删除",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"error": "",
		"msg":   "摧毁成功",
	})

}
