package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"error": "",
		"msg":   msg,
		"data":  data,
	})
}

func Fail(c *gin.Context, code int, err string, msg string, data any) {
	c.JSON(code, gin.H{
		"code":  code,
		"error": err,
		"msg":   msg,
		"data":  data,
	})
}
