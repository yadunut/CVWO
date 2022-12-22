package database

import (
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func InitDB(url string) (DB, error) {

	gormDB, err := gorm.Open(postgres.Open(url))
	if err != nil {
		return DB{}, err
	}

	return DB{gormDB}, nil
}
