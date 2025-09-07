package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {

	var cards []model.Card

	if err := global.DB.Model(&model.Card{}).Where("owner = ?", c.GetString("name")).Find(&cards).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	utils.Ok(c, "查询成功", cards)
}
