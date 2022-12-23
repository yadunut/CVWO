package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID           uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Email        string
	Username     string
	PasswordHash string
}

func NewUser(Email string, Username string, PasswordHash string) *User {
	return &User{
		Email:        Email,
		Username:     Username,
		PasswordHash: PasswordHash,
	}
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
