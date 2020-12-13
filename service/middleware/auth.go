package middleware

import (
	"net/http"

	"template-service/impl/auth"
	i "template-service/interface"

	"github.com/gin-gonic/gin"
)

func AuthByToken(authService i.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUser := authService.AuthorizeUser(ctx)
		if authUser == nil {
			ctx.JSON(http.StatusUnauthorized, "")
			ctx.Abort()

			return
		} else {
			auth.SetAuthUserToContext(ctx, authUser)
		}

		ctx.Next()
	}
}
