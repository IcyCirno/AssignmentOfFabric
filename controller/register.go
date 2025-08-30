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
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "JSON解析失败，请核对填写信息",
		})
		return
	}

	var total int64

	global.DB.Model(&model.User{}).Where("name = ?", iUser.Name).Count(&total)
	if total > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Repeated Username",
			"msg":   "用户名重复",
		})
		return
	}

	global.DB.Model(&model.User{}).Where("email = ?", iUser.Email).Count(&total)
	if total > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Repeated Email",
			"msg":   "邮箱重复",
		})
		return
	}

	pwd, err := utils.Encrypt(iUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "服务器出错！",
		})
		return
	}
	user := model.User{
		Name:     iUser.Name,
		Email:    iUser.Email,
		Password: pwd,
		Gocoin:   viper.GetInt("nft.initasset"),
	}

	if err := global.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "服务器出错！",
		})
		return
	}

	//区块链生成用户

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"error": "",
		"msg":   "注册成功！",
	})

}
