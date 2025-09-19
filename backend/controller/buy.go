package controller

import (
	"blockchain/dto"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BuyRequest 购买卡牌请求参数
// swagger:model BuyRequest
type buy struct {
	// 交易订单ID
	OrderID string `json:"orderid" binding:"required"`
}

// Buy godoc
// @Summary 购买卡牌
// @Description 用户购买市场上的卡牌交易订单，完成金币扣除和卡牌转移
// @Tags NFT
// @Accept json
// @Produce json
// @Param data body buy true "交易订单ID"
// @Success 200 {object} utils.APIResponse[model.CardAndTrans] "交易成功"
// @Failure 400 {object} utils.APIResponse[string] "请求参数错误或资金不足"
// @Failure 500 {object} utils.APIResponse[string] "服务器内部错误或更新失败"
// @Security ApiKeyAuth
// @Router /api/auth/market/buy [post]
func Buy(c *gin.Context) {
	var iBuy buy
	if err := c.ShouldBindJSON(&iBuy); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", "")
		return
	}

	trans, err := dto.GetTransaction(iBuy.OrderID)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", "")
		return
	}

	if trans.Status == "Canceled" {
		utils.Fail(c, http.StatusBadRequest, "", "交易已下架", "")
		return
	}

	user, err := dto.GetUser(trans.Seller)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", "")
		return
	}

	iUser, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "查询出错", "")
		return
	}

	if user.Name == iUser.Name {
		utils.Fail(c, http.StatusBadRequest, "Not", "不能购买自己的卡牌", "")
		return
	}

	if iUser.Gocoin < trans.Price {
		utils.Fail(c, http.StatusBadRequest, "Not Enough", "资金不足", "")
		return
	}

	// 转账给卖家
	user.Gocoin += trans.Price
	for i, id := range user.Cards {
		if id == trans.CardID {
			user.Cards = append(user.Cards[:i], user.Cards[:i+1]...)
			break
		}
	}
	if err := dto.PutUser(user); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", "")
		return
	}

	// 扣除买家金币并添加卡牌
	iUser.Gocoin -= trans.Price
	iUser.Cards = append(iUser.Cards, trans.CardID)
	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", "")
		return
	}

	// 更新交易订单接收者
	trans.Receiver = iUser.Name
	trans.Status = "Done"
	if err := dto.PutTransaction(trans); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", "")
		return
	}

	// 更新卡牌归属
	iCard, err := dto.GetCard(trans.CardID)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值错误", "")
		return
	}

	iCard.OnSale = false
	iCard.TransID = ""
	iCard.Owner = iUser.Name
	if err := dto.PutCard(iCard); err != nil {
		utils.Fail(c, http.StatusInternalServerError, "无法更新", "无法更新", "")
		return
	}

	utils.Ok(c, "交易成功", model.CardAndTrans{
		Card:        iCard,
		Transaction: trans,
	})
}
