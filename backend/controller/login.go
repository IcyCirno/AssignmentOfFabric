package controller

import (
	"blockchain/dto"
	"blockchain/global"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userLogin struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var iUser userLogin

	if err := c.ShouldBindJSON(&iUser); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	user, err := dto.GetUser(iUser.Name)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", nil)
		return
	}

	if !utils.CompareHashAndPassword(user.Password, iUser.Password) {
		utils.Fail(c, http.StatusBadRequest, "Wrong Passwword", "密码错误", nil)
		return
	}

	token, err := utils.GenerateToken(user.Name)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", nil)
		return
	}

	global.RedisClient.Set(user.Name, token)

	utils.Ok(c, "登录成功", token)

}
