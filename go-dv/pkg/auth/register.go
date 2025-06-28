package auth

import (
	"encoding/json"
	"net/http"

	"dv/pkg/errors"
)

var userStore = map[string]string{}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errors.WriteJSONError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	// Check if user already exists
	if _, exists := userStore[creds.Email]; exists {
		errors.WriteJSONError(w, "email already registered", http.StatusConflict)
		return
	}

	// Save user (you can hash password here)
	userStore[creds.Email] = creds.Password

	accessToken, err := GenerateJWT(creds.Email)
	if err != nil {
		errors.WriteJSONError(w, "token generation failed", http.StatusInternalServerError)
		return
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		errors.WriteJSONError(w, "refresh token failed", http.StatusInternalServerError)
		return
	}

	refreshStore[refreshToken] = creds.Email
	setAuthCookies(w, accessToken, refreshToken)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "registration successful",
	})
}
