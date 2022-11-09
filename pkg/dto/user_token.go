package dto

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserRefreshToken struct {
	Token          string
	Username       string
	IsInvalidated  bool
	ExpirationDate time.Time
}

type UserClaims struct {
	jwt.StandardClaims
	UserContext
}

func NewUserClaims(acc *UserContext) *UserClaims {
	return &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
		UserContext: UserContext{
			Username:    acc.Username,
			Email:       acc.Email,
			DisplayName: acc.DisplayName,
		},
	}
}
