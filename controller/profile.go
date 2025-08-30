package controller

import (
	"blockchain/global"
	"blockchain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	var user model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", c.GetString("name")).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "Server Fail",
			"msg":   "服务器出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"error": "",
		"msg":   "成功",
		"data":  user,
	})

}
