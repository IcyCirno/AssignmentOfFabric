package controller

import (
	"blockchain/dto"
	"blockchain/global"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserLogin 登录请求参数
// swagger:model UserLogin
type UserLogin struct {
	// 用户名
	Name string `json:"name" binding:"required"`
	// 密码
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary 用户登录
// @Description 根据用户名和密码进行登录，成功后返回 Token
// @Tags User
// @Accept json
// @Produce json
// @Param data body UserLogin true "登录信息"
// @Success 200 {object} utils.APIResponse[string] "登录成功，返回Token"
// @Failure 400 {object} utils.APIResponse[string] "请求参数错误或密码错误"
// @Failure 500 {object} utils.APIResponse[string] "服务器内部错误"
// @Router /api/login [post]
func Login(c *gin.Context) {
	var iUser UserLogin

	if err := c.ShouldBindJSON(&iUser); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", "")
		return
	}

	user, err := dto.GetUser(iUser.Name)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", "")
		return
	}

	if !utils.CompareHashAndPassword(user.Password, iUser.Password) {
		utils.Fail(c, http.StatusBadRequest, "Wrong Password", "密码错误", "")
		return
	}

	token, err := utils.GenerateToken(user.Name)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", "")
		return
	}

	global.RedisClient.Set(user.Name, token)

	utils.Ok(c, "登录成功", token)
}
