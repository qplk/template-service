package main

import (
	"log"

	"template-service/config"
	"template-service/impl/postgres"

	"template-service/service"
)

func main() {

	_ = postgres.RunMigration(config.Get())

	customService := service.NewCustomerService("0.0.0.0", "8080")

	log.Fatal(customService.Run())
}
