package utils

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/yadunut/CVWO/backend/auth/internal/config"
)

const (
	PG_DUPLICATE = "23505"
)

func IsAlphaNumeric(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	return re.MatchString(s)
}

func ValidateUsername(s string) error {
	if len(s) < 3 {
		return errors.New("username must be greater than 3 characters")
	}
	if len(s) > 10 {
		return errors.New("username must be shorter than 10 characters")
	}
	if !IsAlphaNumeric(s) {
		return fmt.Errorf("%s invalid, only alphabets and numbers allowed", s)
	}
	return nil
}

func ValidatePassword(s string) error {
	if len(s) < 8 {
		return errors.New("password must be longer than 8 characters")
	}
	if len(s) > 30 {
		return errors.New("password must be shorter than 30 characters")
	}

	return nil
}

type jwtClaims struct {
	jwt.StandardClaims
	ID string
}

func GenerateJwtToken(id string, config config.Config) (string, error) {
	claims := jwtClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(config.JwtExpiry)).Unix(),
			Issuer:    "CVWO",
			NotBefore: time.Now().Local().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecret)
}

func ParseJwtToken(encryptedToken string, config config.Config) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(encryptedToken, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return uuid.UUID{}, errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return uuid.UUID{}, errors.New("JWT expired")
	}
	return uuid.Parse(claims.Id)
}
