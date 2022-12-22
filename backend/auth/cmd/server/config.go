package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseUrl string `split_words:"true"`
	Env         string `split_words:"true"`
	Port        string `default:"8080"`
}

func LoadConfig() (c Config, err error) {
	// only call load if .env exists
	if _, err = os.Stat(".env"); !os.IsNotExist(err) {
		err = godotenv.Load()
		if err != nil {
			return
		}

	}
	err = envconfig.Process("CVWO", &c)
	if err != nil {
		return
	}
	return
}
