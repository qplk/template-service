package auth

import (
	"net/http"

	"template-service/models"

	"github.com/gin-gonic/gin"
)

const authUserContextKeyName = "authUser"

func getTokenFromRequest(req *http.Request) string {
	var token string

	token = req.Header.Get("Authorization")

	if token != "" {
		return token
	}

	cookie, err := req.Cookie("Authorization")

	if err != nil || cookie == nil {
		return ""
	}

	token = cookie.Value

	return token
}

func SetAuthUserToContext(ctx *gin.Context, authUser *models.User) {
	ctx.Set(authUserContextKeyName, authUser)
}

func GetAuthUserFromContext(ctx *gin.Context) *models.User {
	data, exists := ctx.Get(authUserContextKeyName)

	if !exists {
		return nil
	}

	authUser, ok := data.(models.User)
	if !ok {
		return nil
	}

	return &authUser
}
