package _interface

import (
	"template-service/models"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	AuthorizeUser(ctx *gin.Context) *models.User
}
