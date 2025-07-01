package auth

import (
	"encoding/json"
	"net/http"

	"dv/pkg/errors"
)

func (auth *Auth) Login(w http.ResponseWriter, r *http.Request) {
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

	user, exists := (*auth.UserService).Exists(creds.Email)
	if !exists {
		errors.WriteJSONError(w, "user not found", http.StatusNotFound)
		return
	}

	if user == nil {
		errors.WriteJSONError(w, "user not found", http.StatusNotFound)
		return
	}

	if user.Password != creds.Password {
		errors.WriteJSONError(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, err := auth.GenerateJWT(user.Email)
	if err != nil {
		errors.WriteJSONError(w, "token generation failed", http.StatusInternalServerError)
		return
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		errors.WriteJSONError(w, "refresh token failed", http.StatusInternalServerError)
		return
	}

	auth.RefreshStore[refreshToken] = user.Email
	auth.SetAuthCookies(w, accessToken, refreshToken)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "login successful",
	})
}
