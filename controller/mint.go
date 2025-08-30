package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cardCreate struct {
	Name    string `json:"name" binding:"required"`
	Profile string `json:"profile" binding:"required"`
	Data    string `json:"data" binding:"required"`
}

func Mint(c *gin.Context) {

	var info cardCreate
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "JSON解析失败，请核对填写信息",
		})
		return
	}
	owner := c.GetString("name")
	var iUser model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", owner).First(&iUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "服务器出错",
		})
		return
	}

	if iUser.Gocoin < 10 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "",
			"msg":   "资金不足",
		})
		return
	}

	iCard := model.Card{
		Name:    info.Name,
		Profile: info.Profile,
		HashID:  utils.GenerateCardID(info.Name, info.Profile, owner),
		Owner:   owner,
		Avatar:  info.Data,

		Attack: utils.RandomInt(0, 99),
		Blood:  utils.RandomInt(0, 99),
		Cost:   utils.RandomInt(0, 10),
		Rarity: utils.RandomRarity(),

		OnSale: false,
	}

	//先上链

	if err := global.DB.Save(&iCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "服务器错误",
		})
		return
	}

	if err := global.DB.Model(&iUser).Update("gocoin", iUser.Gocoin-10).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "服务器错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"error": "",
		"msg":   "铸造成功",
		"data":  iCard,
	})

}
