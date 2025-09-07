package controller

import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type hashID struct {
	HashID string `json:"hashid" binding:"required"`
}

func Destroy(c *gin.Context) {
	var ihash hashID
	if err := c.ShouldBindJSON(&ihash); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	var iCard model.Card
	if err := global.DB.Model(&model.Card{}).Where("hash_id = ?", ihash.HashID).First(&iCard).Error; err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "哈希值错误", nil)
		return
	}

	if iCard.OnSale {
		utils.Fail(c, http.StatusBadRequest, "", "卡牌正在市场中", nil)
		return
	}

	var iUser model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", iCard.Owner).First(&iUser).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "无法找到", nil)
		return
	}

	if err := global.DB.Model(&iUser).Update("gocoin", iUser.Gocoin+4).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "无法找到", nil)
		return
	}

	if err := global.DB.Delete(&iCard).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "删除失败", nil)
		return
	}

	utils.Ok(c, "摧毁成功", nil)

}
