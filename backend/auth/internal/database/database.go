package database

import (
	"github.com/yadunut/CVWO/backend/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewUser(Email string, Username string, PasswordHash string) models.User {
	return models.User{
		Email:        Email,
		Username:     Username,
		PasswordHash: PasswordHash,
	}
}

type DB struct {
	*gorm.DB
}

func Init(url string, log logger.Interface) (DB, error) {
	if log == nil {
		log = logger.Default
	}

	gormDB, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		return DB{}, err
	}

	return DB{gormDB}, nil
}
