package profile

import (
	"template-service/config"
	"template-service/impl/person"
)

type Handler struct {
	cfg           *config.Config
	personService *person.PersonInfoService
}

func NewHandler(
	cfg *config.Config,
	personService *person.PersonInfoService,
) *Handler {
	return &Handler{
		cfg:           cfg,
		personService: personService,
	}
}
