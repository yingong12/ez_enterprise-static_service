package controller

import "github.com/gin-gonic/gin"

func Healthy(ctx *gin.Context) {
	ctx.Writer.Write([]byte("ok"))
}
