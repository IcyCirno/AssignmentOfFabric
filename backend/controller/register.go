package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type registerUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var iUser registerUser
	if err := c.ShouldBindJSON(&iUser); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	var total int64

	global.DB.Model(&model.User{}).Where("name = ?", iUser.Name).Count(&total)
	if total > 0 {
		utils.Fail(c, http.StatusBadRequest, "Repeated Username", "用户名重复", nil)
		return
	}

	global.DB.Model(&model.User{}).Where("email = ?", iUser.Email).Count(&total)
	if total > 0 {
		utils.Fail(c, http.StatusBadRequest, "Repeated Email", "邮箱重复", nil)
		return
	}

	pwd, err := utils.Encrypt(iUser.Password)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "Server Fail", "服务器出错", nil)
		return
	}
	user := model.User{
		Name:     iUser.Name,
		Email:    iUser.Email,
		Password: pwd,
		Rank:     0,
		Gocoin:   viper.GetInt("nft.initasset"),
		Mine:     false,
	}

	if err := global.DB.Save(&user).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "Server Fail", "服务器出错", nil)
		return
	}

	//区块链生成用户

	utils.Ok(c, "注册成功", nil)

}
