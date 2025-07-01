package auth

import (
	"encoding/json"
	"net/http"

	"dv/internal/users"
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

	userCreds := users.UserDTO{
		Email:    creds.Email,
		Password: creds.Password,
	}

	user, err := auth.UserService.Login(userCreds)
	if err != nil {
		if err.Error() == "invalid credentials or user inactive" {
			errors.WriteJSONError(w, "invalid credentials or user inactive", http.StatusUnauthorized)
			return
		}
		errors.WriteJSONError(w, "login failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !user.Status {
		errors.WriteJSONError(w, "user is inactive", http.StatusForbidden)
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
