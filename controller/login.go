package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var iUser userLogin

	if err := c.ShouldBindJSON(&iUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "JSON解析失败，请核对填写信息",
		})
		return
	}

	var user model.User
	err := global.DB.Model(&model.User{}).Where("email = ?", iUser.Email).Find(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "邮箱未注册!",
		})
		return
	}

	if !utils.CompareHashAndPassword(user.Password, iUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Wrong Passwword",
			"msg":   "密码错误！",
		})
		return
	}

	token, err := utils.GenerateToken(user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "服务器出错！",
		})
		return
	}

	global.RedisClient.Set(user.Name, token)

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"error": "",
		"msg":   "登录成功",
		"token": token,
	})

}
