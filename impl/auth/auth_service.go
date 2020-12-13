package auth

import (
	"template-service/impl/repository"
	"template-service/models"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	TokenRepo *repository.TokenRepo
}

func NewAuthService(
	tokenRepo *repository.TokenRepo,
) *AuthService {
	return &AuthService{
		TokenRepo: tokenRepo,
	}
}

func (as *AuthService) AuthorizeUser(ctx *gin.Context) *models.User {

	token := getTokenFromRequest(ctx.Request)

	if token == "" {
		return nil
	}

	authUser, err := as.TokenRepo.GetUserByToken(token)

	if err != nil || authUser == nil {
		return nil
	}

	return authUser
}
