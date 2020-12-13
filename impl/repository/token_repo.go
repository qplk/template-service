package repository

import (
	"github.com/jmoiron/sqlx"
	"template-service/config"
	"template-service/models"
)

type TokenRepo struct {
	cfg  *config.Config
	conn *sqlx.DB
}

func NewTokenRepo(
	cfg *config.Config,
	conn *sqlx.DB,
) *TokenRepo {
	return &TokenRepo{
		cfg:  cfg,
		conn: conn,
	}
}

func (tr *TokenRepo) GetUserByToken(token string) (*models.User, error) {

	//todo: return user information by given token
	return &models.User{}, nil
}
