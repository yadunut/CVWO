package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port           string `default:"8080"`
	AuthServiceUrl string `split_words:"true"`
}

func LoadConfig() (c Config, err error) {
	// only call load if .env exists
	if _, err = os.Stat(".env"); !os.IsNotExist(err) {
		err = godotenv.Load()
		if err != nil {
			return
		}

	}

	if err != nil {
		return
	}

	err = envconfig.Process("CVWO", &c)
	if err != nil {
		return
	}
	return
}
