package middleware

import (
	"blockchain/global"
	"blockchain/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {

		token := c.GetHeader(global.TOKEN_NAME)
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": "Empty Token",
				"msg":   "Token为空",
			})
			return
		}

		if !strings.HasPrefix(token, global.TOKEN_PREFIX) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": "Wrong Prefix",
				"msg":   "前缀出错！",
			})
			return
		}

		token = token[len(global.TOKEN_PREFIX):]
		iJwt, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": "Bad Token",
				"msg":   "Token非法或错误！",
			})
			return
		}

		t, err := global.RedisClient.Get(iJwt.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  http.StatusInternalServerError,
				"error": err.Error(),
				"msg":   "对比Token失败！",
			})
			return
		}

		if t != token {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": "Bad Token",
				"msg":   "Token过期！",
			})
			return
		}

		c.Set("name", iJwt.Name)

		c.Next()

	}
}
