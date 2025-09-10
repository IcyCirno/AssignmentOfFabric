package controller

import (
	"blockchain/dto"
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

	trans, err := dto.GetTransaction(iBuy.OrderID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	user, err := dto.GetUser(trans.Receiver)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	iUser, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", nil)
		return
	}

	if user.Name == iUser.Name {
		utils.Fail(c, http.StatusBadRequest, "Not", "Not", nil)
	}

	if iUser.Gocoin < trans.Price {
		utils.Fail(c, http.StatusBadRequest, "Not Enough", "资金不足", nil)
		return
	}

	user.Gocoin += trans.Price
	for i, id := range user.Cards {
		if id == trans.CardID {
			user.Cards = append(user.Cards[:i], user.Cards[:i+1]...)
			break
		}
	}
	if err := dto.PutUser(user); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	iUser.Gocoin -= trans.Price
	iUser.Cards = append(iUser.Cards, trans.CardID)
	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	trans.Receiver = iUser.Name
	if err := dto.PutTransaction(trans); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", nil)
		return
	}

	iCard, err := dto.GetCard(trans.CardID)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值错误", nil)
		return
	}

	iCard.OnSale = false
	iCard.Owner = iUser.Name
	if err := dto.PutCard(iCard); err != nil {
		utils.Fail(c, http.StatusInternalServerError, "无法更新", "无法更新", nil)
		return
	}

	utils.Ok(c, "交易成功", nil)
}
