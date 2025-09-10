package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagger:model APIResponse
type APIResponse[T any] struct {
	Code  int    `json:"code" example:"200"` // 状态码
	Error string `json:"error" example:""`   // 错误信息
	Msg   string `json:"msg" example:"请求成功"` // 描述信息
	Data  T      `json:"data"`               // 返回数据，可以是任何类型
}

func Ok[T any](c *gin.Context, msg string, data T) {
	c.JSON(http.StatusOK, APIResponse[T]{
		Code:  http.StatusOK,
		Error: "",
		Msg:   msg,
		Data:  data,
	})
}

func Fail[T any](c *gin.Context, code int, err string, msg string, data T) {
	c.JSON(code, APIResponse[T]{
		Code:  code,
		Error: err,
		Msg:   msg,
		Data:  data,
	})
}

/*
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
*/
