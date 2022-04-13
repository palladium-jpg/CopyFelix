package middleware

import (
	"CopyFelix/common"
	"CopyFelix/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取authorization header
		tokenString := context.GetHeader("Authorization")
		//获取头文件
		if len(tokenString) == 0 || !strings.HasPrefix(tokenString, "Bearer") {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			context.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			context.Abort()
			return
		}

		userId := claims.UserId
		DB := common.InitDb()
		var user model.User
		user.ID = userId
		get, erG := DB.Get(&user)

		if erG != nil {
			return
		}

		if !get {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			context.Abort()
			return
		}

		context.Set("user", user)
		context.Next()
	}
}
