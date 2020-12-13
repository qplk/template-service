package person

import (
	"template-service/config"
	"template-service/impl/repository"
)

type PersonInfoService struct {
	cfg      *config.Config
	userRepo *repository.UserRepo
}

func NewPersonInfoService(
	cfg *config.Config,
	userRepo *repository.UserRepo,
) *PersonInfoService {
	return &PersonInfoService{
		cfg:      cfg,
		userRepo: userRepo,
	}
}
