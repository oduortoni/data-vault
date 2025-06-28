package auth

import (
	"encoding/json"
	"net/http"

	"dv/pkg/errors"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		errors.WriteJSONError(w, "invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if creds.Email == "" || creds.Password == "" {
		errors.WriteJSONError(w, "email and password required", http.StatusBadRequest)
		return
	}

	user, err := FindUserByEmail(creds.Email)
	if err != nil || user.Password != creds.Password {
		errors.WriteJSONError(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, err := GenerateJWT(user.Email)
	if err != nil {
		errors.WriteJSONError(w, "token generation failed", http.StatusInternalServerError)
		return
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		errors.WriteJSONError(w, "refresh token failed", http.StatusInternalServerError)
		return
	}

	refreshStore[refreshToken] = user.Email
	setAuthCookies(w, accessToken, refreshToken)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "login successful",
	})
}
