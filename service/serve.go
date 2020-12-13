package service

import (
	"template-service/config"
	"template-service/impl/auth"
	"template-service/impl/person"
	"template-service/impl/postgres"
	"template-service/impl/repository"
	"template-service/service/middleware"
	"template-service/service/profile"

	"github.com/gin-gonic/gin"
)

type CustomService struct {
	Host  string
	Port  string
	Route *gin.Engine
}

func NewCustomerService(
	host string,
	port string,
) *CustomService {
	return &CustomService{
		Host:  host,
		Port:  port,
		Route: gin.New(),
	}
}

func (cs *CustomService) Run() error {
	cfg := config.Get()

	route := cs.Route
	route.Use(gin.Recovery())

	cs.SetupServiceRoutes(cfg)

	return cs.Route.Run(cs.Host + ":" + cs.Port)
}

func (cs *CustomService) SetupServiceRoutes(cfg *config.Config) {

	pgClient := postgres.GetClient(cfg)
	tokenRepo := repository.NewTokenRepo(cfg, pgClient)

	userRepo := repository.NewUserRepo(cfg, pgClient)
	personService := person.NewPersonInfoService(cfg, userRepo)
	profileHandler := profile.NewHandler(cfg, personService)

	//authorized routes
	authR := cs.Route.Group("/profile")
	{
		authService := auth.NewAuthService(tokenRepo)
		authR.Use(middleware.AuthByToken(authService))

		authR.GET("/orders", profileHandler.UserOrdersHandler)
	}

	//unauthorized routes
	noAuthR := cs.Route.Group("")
	{
		noAuthR.GET("/login")
	}

	return
}
