package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// DestroyRequest 摧毁卡牌请求参数
// swagger:model DestroyRequest
type hashID struct {
	// 卡牌唯一ID
	HashID string `json:"hashid" binding:"required"`
}

// Destroy godoc
// @Summary 摧毁卡牌
// @Description 用户摧毁自己拥有的卡牌，卡牌不在市场中才可摧毁，摧毁后返还部分金币
// @Tags NFT
// @Accept json
// @Produce json
// @Param data body hashID true "卡牌哈希ID"
// @Success 200 {object} utils.APIResponse[string] "摧毁成功"
// @Failure 400 {object} utils.APIResponse[string] "请求参数错误或卡牌在市场中"
// @Failure 500 {object} utils.APIResponse[string] "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/auth/card/destroy [post]
func Destroy(c *gin.Context) {
	var ihash hashID
	if err := c.ShouldBindJSON(&ihash); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", "")
		return
	}

	iCard, err := dto.GetCard(ihash.HashID)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值错误", "")
		return
	}

	if iCard.OnSale {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌正在市场中", "")
		return
	}

	iUser, err := dto.GetUser(iCard.Owner)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "无法找到", "")
		return
	}

	iUser.Gocoin += viper.GetInt("nft.destroy")
	for i, id := range iUser.Cards {
		if id == iCard.HashID {
			iUser.Cards = append(iUser.Cards[:i], iUser.Cards[i+1:]...)
			break
		}
	}
	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "无法找到", "")
		return
	}

	iCard.Destroy = true
	if err := dto.PutCard(iCard); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "删除失败", "")
		return
	}

	utils.Ok(c, "摧毁成功", "")
}
