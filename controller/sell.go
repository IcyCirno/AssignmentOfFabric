package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sell struct {
	HashID string `json:"hashid" binding:"required"`
	Cost   int    `json:"cost" binding:"required"`
}

func Sell(c *gin.Context) {
	var isell sell
	if err := c.ShouldBindJSON(&isell); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	if isell.Cost < 0 {
		utils.Fail(c, http.StatusBadRequest, "", "交易金额不合法", nil)
		return
	}

	var iCard model.Card
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", isell.HashID).First(&iCard).Error; err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值出错", nil)
		return
	}

	if iCard.OnSale {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌正在市场中", nil)
		return
	}

	trans := model.Transaction{
		CardID:  iCard.HashID,
		Seller:  iCard.Owner,
		TransID: utils.GenerateOrderID(),
		Price:   isell.Cost,
	}

	if err := global.DB.Model(&iCard).Update("on_sale", true).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "无法更新", "无法更新", nil)
		return
	}

	if err := global.DB.Save(&trans).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "无法创建订单", "无法创建订单", nil)
		return
	}

	utils.Ok(c, "创建交易成功", trans)

}
