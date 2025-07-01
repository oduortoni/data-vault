package auth

import (
	"dv/internal/users"
)

type Auth struct {
	UserService  *users.UserService
	JWTSecret    []byte
	TokenName    string
	RefreshStore map[string]string
}

// NewAuth creates a new Auth instance
func NewAuth(userService *users.UserService, jwtSecret []byte, tokenName string) *Auth {
	return &Auth{
		UserService:  userService,
		JWTSecret:    jwtSecret,
		TokenName:    tokenName,
		RefreshStore: map[string]string{}, // refreshToken -> email
	}
}
