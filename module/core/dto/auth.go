package dto

import (
	"fx-golang-server/pkg/e"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Phone string `json:"phone"`
}

type CreateUserRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type CreateUserResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserPayload struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (c UserPayload) Valid() error {
	if c.ExpiresAt < time.Now().Unix() {
		return e.ErrTokenExpired
	}
	return nil
}
