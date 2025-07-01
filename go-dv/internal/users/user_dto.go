package users

import (
	"errors"
	"strings"
)

type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

func (u *UserDTO) Validate() error {
	if u.Username == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if !isValidEmail(u.Email) {
		return errors.New("invalid email format")
	}
	return nil
}

func isValidEmail(email string) bool {
	if len(email) < 3 || !strings.Contains(email, "@") {
		return false
	}
	return true
}
