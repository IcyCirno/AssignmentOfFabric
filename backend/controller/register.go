package controller

import (
	"blockchain/dto"
	"blockchain/fabric"
	"blockchain/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RegisterUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var iUser RegisterUser
	if err := c.ShouldBindJSON(&iUser); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", "")
		return
	}

	data, err := fabric.Contract.EvaluateTransaction("GetUser", iUser.Name)
	if data != nil {
		utils.Fail(c, http.StatusInternalServerError, "", "用户存在", "")
		return
	}

	pwd, err := utils.Encrypt(iUser.Password)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "Server Fail", "服务器出错", "")
		return
	}

	user := dto.User{
		Name:     iUser.Name,
		Password: pwd,
		CreateAt: time.Now(),
		Rank:     0,
		Gocoin:   viper.GetInt("nft.initasset"),
		EndTime:  time.Now(),
	}

	for i := 1; i <= 3; i++ {
		card, err := utils.CreateCard("init_"+strconv.Itoa(i), iUser.Name, 0)
		if err != nil {
			utils.Fail(c, http.StatusInternalServerError, err.Error(), "Fail to Init User", "")
			return
		}
		user.Cards = append(user.Cards, card.HashID)
		if err := dto.PutCard(card); err != nil {
			utils.Fail(c, http.StatusInternalServerError, err.Error(), "添加卡牌出错", "")
			return
		}
	}

	if err := dto.PutUser(user); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "Fabric Fail", "")
		return
	}

	utils.Ok(c, "注册成功", "")
}
