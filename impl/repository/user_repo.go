package repository

import (
	"time"

	"template-service/config"
	"template-service/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	cfg  *config.Config
	conn *sqlx.DB
}

func NewUserRepo(
	cfg *config.Config,
	conn *sqlx.DB,
) *UserRepo {
	return &UserRepo{
		cfg:  cfg,
		conn: conn,
	}
}

//todo: choose user identifier (login/id/token)
func (ur *UserRepo) GetOrdersByLogin(login string) []models.Order {

	orders := make([]models.Order, 0)
	orders = append(orders, models.Order{
		Type:     "ТО",
		Cost:     1000,
		Time:     time.Now(),
		Duration: 60,
	})

	return orders
}
