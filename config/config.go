package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v6"
)

var (
	config *Config
	mux    sync.Mutex
)

type Config struct {
	Postgres Postgres
}

func Get() *Config {
	if config == nil {
		initConfig()
	}

	return config
}

func initConfig() {
	mux.Lock()
	defer mux.Unlock()

	var conf Config
	if err := env.Parse(&conf); err != nil {
		log.Println("Can not parse env")
	}

	config = &conf
}
