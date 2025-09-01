package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Market(c *gin.Context) {
	var transactions []model.Transaction
	if err := global.DB.Model(&model.Transaction{}).Find(&transactions).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	utils.Ok(c, "查询成功", transactions)

}
