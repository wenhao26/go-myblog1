package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SuccessCode = 0
	FailedCode  = -1

	// TODO 更多状态码
)

type Base struct {
}

var BaseAPI = &Base{}

func (*Base) Default(ctx *gin.Context) {
	ctx.String(http.StatusForbidden, "Forbidden")
	//ctx.AbortWithStatus(402)
}

func (*Base) Success(ctx *gin.Context, msg string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": SuccessCode,
		"msg":  msg,
		"data": data,
	})
}

func (*Base) Failed(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": gin.H{},
	})
}
