package controller

/*
import (
	"blockchain/global"
	"blockchain/model"
	"blockchain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type attack struct {
	A string `json:"a" binding:"required"`
	B string `json:"b" binding:"required"`
	C string `json:"c" binding:"required"`
}

func Attack(c *gin.Context) {
	var info attack
	if err := c.ShouldBind(&info); err != nil {
		utils.Fail(c, http.StatusBadRequest, err.Error(), "JSON解析失败，请核对填写信息", nil)
		return
	}

	var user model.User
	for {
		if err := global.DB.Order("RAND()").Limit(1).First(&user).Error; err != nil {
			utils.Fail(c, http.StatusInternalServerError, err.Error(), "数据库出错", nil)
			return
		}
		if user.Name != c.GetString("name") {
			break
		}
	}

	var iUser model.User
	if err := global.DB.Model(&model.User{}).Where("name = ?", c.GetString("name")).First(&iUser).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "无法找到", nil)
		return
	}

	rank := 5
	ok := Fight(info, user)
	iUser.Rank += rank * ok
	user.Rank += rank * -ok

	if iUser.Rank < 0 {
		iUser.Rank = 0
	}
	if user.Rank < 0 {
		user.Rank = 0
	}

	if err := global.DB.Save(&iUser).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错", nil)
		return
	}
	if err := global.DB.Save(&user).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, err.Error(), "服务器出错", nil)
		return
	}

	if ok == 1 {
		utils.Ok(c, "胜利", rank)
	} else {
		utils.Ok(c, "失败", rank)
	}

}

func Fight(info attack, user model.User) int {
	return 1
}
*/
