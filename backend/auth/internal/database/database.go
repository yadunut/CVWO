package database

import (
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
	email        string
	Username     string
	PasswordSalt string
	PasswordHash string
}

type DB struct {
	*gorm.DB
}

func InitDB(url string, log logger.Interface) (DB, error) {
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
