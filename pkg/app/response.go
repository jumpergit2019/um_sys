package app

import (
	"um_sys/pkg/e"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (this *Gin) Response(httpCode int, errCode int, data interface{}) {
	this.Ctx.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
