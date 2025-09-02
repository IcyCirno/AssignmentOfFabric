package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type buy struct {
	OrderID string `json:"orderid" binding:"required"`
}

func Buy(c *gin.Context) {
	var iBuy buy
	if err := c.ShouldBindJSON(&iBuy); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	var trans model.Transaction
	if err := global.DB.Model(&model.Transaction{}).Where("trans_id = ?", iBuy.OrderID).First(&trans).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	var user model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", trans.Seller).First(&user).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	var iUser model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", c.GetString("name")).First(&iUser).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	if user.Name == iUser.Name {
		utils.Fail(c, http.StatusBadRequest, "Not", "Not", nil)
	}

	// if iUser.Gocoin < trans.Price {
	// 	utils.Fail(c, http.StatusBadRequest, "Not Enough", "资金不足", nil)
	// 	return
	// }

	if err := global.DB.Model(&user).Update("gocoin", user.Gocoin+trans.Price).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	if err := global.DB.Model(&iUser).Update("gocoin", iUser.Gocoin-trans.Price).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	if err := global.DB.Model(&trans).Update("Receiver", iUser.Name).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	var iCard model.Card
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", trans.CardID).First(&iCard).Error; err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值错误", nil)
		return
	}

	if err := global.DB.Model(&iCard).Update("on_sale", false).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "无法更新", "无法更新", nil)
		return
	}

	if err := global.DB.Model(&iCard).Update("owner", iUser.Name).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "无法更新", "无法更新", nil)
		return
	}

	utils.Ok(c, "交易成功", nil)
}
