package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type defense struct {
	A string `json:"a" binding:"required"`
	B string `json:"b" binding:"required"`
	C string `json:"c" binding:"required"`
}

func Defense(c *gin.Context) {
	var info defense
	if err := c.ShouldBind(&info); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	var x, y, z model.Card
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", info.A).First(&x).Error; err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值出错", nil)
		return
	}
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", info.B).First(&y).Error; err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值出错", nil)
		return
	}
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", info.C).First(&z).Error; err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值出错", nil)
		return
	}

	if x.OnSale || y.OnSale || z.OnSale {
		utils.Fail(c, http.StatusBadRequest, "", "有卡牌正在市场！", nil)
		return
	}

	if x.HashID == y.HashID || x.HashID == z.HashID || y.HashID == z.HashID {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌重复!", nil)
		return
	}

	var iUser model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", c.GetString("name")).First(&iUser).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "无法找到", nil)
		return
	}

	iUser.A = info.A
	iUser.B = info.B
	iUser.C = info.C

	if err := global.DB.Save(&iUser).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错", nil)
	}

	utils.Ok(c, "设置成功", nil)

}
