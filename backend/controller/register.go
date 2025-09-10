package controller

import (
	"blockchain/dto"
	"blockchain/fabric"
	"blockchain/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type registerUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var iUser registerUser
	if err := c.ShouldBindJSON(&iUser); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	data, err := fabric.Contract.EvaluateTransaction("GetUser", iUser.Name)
	/*if err != nil && err.Error() != "Not Found" {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "Fabric Fail", nil)
		return
	}*/
	if data != nil {
		utils.Fail(c, http.StatusInternalServerError, "", "用户存在", nil)
		return
	}

	pwd, err := utils.Encrypt(iUser.Password)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "Server Fail", "服务器出错", nil)
		return
	}

	user := dto.User{
		Name:     iUser.Name,
		Password: pwd,
		CreateAt: time.Now(),

		Rank:   0,
		Gocoin: viper.GetInt("nft.initasset"),

		EndTime: time.Now(),
	}

	if err := dto.PutUser(user); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "Fabric Fail", nil)
		return
	}

	utils.Ok(c, "注册成功", nil)

}
