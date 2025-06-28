package auth

import (
	"errors"
	"net/http"
	"time"
)

var (
	jwtSecret    = []byte("super-secret-key")
	tokenName    = "auth_token"
	refreshStore = map[string]string{} // refreshToken -> email
)

type User struct {
	Email    string
	Password string
}

func FindUserByEmail(email string) (*User, error) {
	if email == "admin@example.com" {
		return &User{Email: email, Password: "password123"}, nil
	}
	return nil, errors.New("user not found")
}

func setAuthCookies(w http.ResponseWriter, accessToken, refreshToken string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(15 * time.Minute),
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/auth/refresh",
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})
}
