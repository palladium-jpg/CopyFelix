package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, h gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, h, msg)
}
func Fail(ctx *gin.Context, h gin.H, string2 string) {
	Response(ctx, http.StatusBadRequest, 400, h, string2)
}
