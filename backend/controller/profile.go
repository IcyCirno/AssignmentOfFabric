package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Profile godoc
// @Summary 获取用户信息
// @Description 根据当前登录用户的名字获取用户详细信息
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} utils.APIResponse[dto.User] "请求成功，返回用户信息"
// @Failure 500 {object} utils.APIResponse[string] "服务器错误"
// @Security ApiKeyAuth
// @Router /api/auth/user/profile [post]
func Profile(c *gin.Context) {
	user, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", "")
		return
	}

	utils.Ok(c, "成功", user)
}
