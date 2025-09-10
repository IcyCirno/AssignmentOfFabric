package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Market(c *gin.Context) {
	var transactions []dto.Transaction

	iUser, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	for _, id := range iUser.Trans {
		trans, err := dto.GetTransaction(id)
		if err != nil {
			utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
			return
		}
		transactions = append(transactions, trans)
	}

	utils.Ok(c, "查询成功", transactions)

}
