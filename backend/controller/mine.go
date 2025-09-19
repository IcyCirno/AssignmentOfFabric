package controller

import (
	"blockchain/dto"
	"blockchain/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// mineRequest 挖矿请求参数
// swagger:model mineRequest
type mine struct {
	Difficulty string `json:"difficulty" binding:"required"`
	A          string `json:"a" binding:"required"`
	B          string `json:"b" binding:"required"`
	C          string `json:"c" binding:"required"`
}

// Mine godoc
// @Summary 挖矿接口
// @Description 用户对指定卡牌进行挖矿操作，冷却中或卡牌在市场上无法挖矿
// @Tags NFT
// @Accept json
// @Produce json
// @Param data body mine true "挖矿信息"
// @Success 200 {object} utils.APIResponse[string] "挖矿成功"
// @Failure 400 {object} utils.APIResponse[string] "请求参数错误或冷却中/卡牌在市场上"
// @Failure 500 {object} utils.APIResponse[string] "服务器内部错误"
// @Security ApiKeyAuth
// @Router /api/auth/user/mine [post]
func Mine(c *gin.Context) {
	var info mine

	if err := c.ShouldBind(&info); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", "")
		return
	}

	if info.A == info.B || info.A == info.C || info.B == info.C {
		utils.Fail(c, http.StatusBadRequest, "", "不能重复！", "")
		return
	}

	iUser, err := dto.GetUser(c.GetString("name"))
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", "")
		return
	}

	if time.Now().Before(iUser.EndTime.Add(viper.GetDuration("nft.minetime"))) {
		utils.Fail(c, http.StatusBadRequest, "", "冷却中", "")
		return
	}

	iUser.EndTime = time.Now().Add(viper.GetDuration("nft.minetime"))

	A, err := dto.GetCard(info.A)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", "")
		return
	}
	if !A.Free() {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌不在空闲状态", "")
		return
	}

	B, err := dto.GetCard(info.B)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", "")
		return
	}
	if !B.Free() {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌不在空闲状态", "")
		return
	}

	C, err := dto.GetCard(info.C)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错！", "")
		return
	}
	if !C.Free() {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌不在空闲状态", "")
		return
	}

	ok, money := utils.RandomMine(info.Difficulty, A, B, C)
	if !ok {
		utils.Ok(c, "挑战失败！", "")
		return
	}

	iUser.Gocoin += money
	if err := dto.PutUser(iUser); err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "更新失败", "")
		return
	}

	utils.Ok(c, "挖矿成功", "")
}
