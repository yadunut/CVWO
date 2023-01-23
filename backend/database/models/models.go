package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type User struct {
	Model        `gorm:"embedded"`
	Email        string
	Username     string
	PasswordHash string
	Threads      []Thread  `gorm:"foreignKey:OwnerId"`
	Comments     []Comment `gorm:"foreignKey:OwnerId"`
}

func NewUser(Email string, Username string, PasswordHash string) User {
	return User{
		Email:        Email,
		Username:     Username,
		PasswordHash: PasswordHash,
	}
}

type Thread struct {
	Model    `gorm:"embedded"`
	OwnerId  uuid.UUID
	Title    string
	Body     string
	Comments []Comment `gorm:"foreignKey:ThreadId"`
}

func NewThread(OwnerId uuid.UUID, Title string, Body string) Thread {
	return Thread{
		OwnerId: OwnerId,
		Title:   Title,
		Body:    Body,
	}
}

type Comment struct {
	Model    `gorm:"embedded"`
	OwnerId  uuid.UUID
	ThreadId uuid.UUID
	Body     string
}
