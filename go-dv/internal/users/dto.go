package users

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
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

	if u.Password == "" {
		return errors.New("password is required")
	}

	if !isValidPassword(u.Password) {
		return errors.New("password must be at least 6 characters long")
	}

	if u.Status != true && u.Status != false {
		return errors.New("status must be true or false")
	}

	if len(u.Username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	if len(u.Email) < 3 {
		return errors.New("email must be at least 3 characters long")
	}
	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	if strings.Contains(u.Username, " ") {
		return errors.New("username cannot contain spaces")
	}
	if strings.Contains(u.Email, " ") {
		return errors.New("email cannot contain spaces")
	}
	if strings.Contains(u.Password, " ") {
		return errors.New("password cannot contain spaces")
	}
	if strings.Contains(u.Username, "@") {
		return errors.New("username cannot contain '@' character")
	}
	if strings.Contains(u.Email, "#") {
		return errors.New("email cannot contain '#' character")
	}
	if strings.Contains(u.Password, "#") {
		return errors.New("password cannot contain '#' character")
	}
	if strings.Contains(u.Username, "#") {
		return errors.New("username cannot contain '#' character")
	}
	if strings.Contains(u.Email, "$") {
		return errors.New("email cannot contain '$' character")
	}
	if strings.Contains(u.Password, "$") {
		return errors.New("password cannot contain '$' character")
	}
	if strings.Contains(u.Username, "$") {
		return errors.New("username cannot contain '$' character")
	}
	return nil
}

func isValidEmail(email string) bool {
	if len(email) < 3 || !strings.Contains(email, "@") {
		return false
	}
	return true
}

// TODO: check for the strength
func isValidPassword(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true
}

func (u UserDTO) HashPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u UserDTO) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *UserDTO) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s", u.ID, u.Username, u.Email)
}
