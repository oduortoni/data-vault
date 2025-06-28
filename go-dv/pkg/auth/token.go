package auth

import (
	"crypto/rand"
	"encoding/base64"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

// Random string for refresh token
func generateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
