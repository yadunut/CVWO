package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseUrl string `split_words:"true"`
	Port        string `default:"8080"`
	Host        string `default:"0.0.0.0"`
}

func Load() (c Config, err error) {
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
	cRef := reflect.ValueOf(&c).Elem()
	for i := 0; i < cRef.NumField(); i++ {
		field := cRef.Field(i)
		if field.IsZero() {
			err = fmt.Errorf("%s cannot be empty", cRef.Type().Field(i).Name)
			return
		}
	}
	return
}
