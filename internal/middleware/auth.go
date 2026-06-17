package middleware

import (
	"strings"
	"user/pkg/ecode"
	auth "user/pkg/jwt"
	"user/pkg/response"

	"github.com/gin-gonic/gin"
)

const CurrentUserIDKey = "current_user_id"
const CurrentUsername = "current_user_name"

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			response.Fail(ctx, ecode.TokenInvalid, "缺少 Authorization 请求头")
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Fail(ctx, ecode.TokenInvalid, "Authorization 格式错误")
			ctx.Abort()
			return
		}
		tokenString := parts[1]
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			response.Fail(ctx, ecode.TokenInvalid, "Token 无效或已过期")
			ctx.Abort()
			return
		}
		ctx.Set(CurrentUserIDKey, claims.UserID)
		ctx.Set(CurrentUsername, claims.Username)

		ctx.Next()
	}
}
