package dto

import (
	"fx-golang-server/pkg/e"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginRequest struct {
	Phone string `json:"phone"`
}

type CreateUserRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type UserPayload struct {
	jwt.StandardClaims
}

func (c UserPayload) Valid() error {
	if c.ExpiresAt < time.Now().Unix() {
		return e.ErrTokenExpired
	}
	return nil
}
