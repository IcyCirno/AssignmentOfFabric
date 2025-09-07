package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	var user model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", c.GetString("name")).Find(&user).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "Server Fail", "服务器出错", nil)
		return
	}

	utils.Ok(c, "成功", user)

}
