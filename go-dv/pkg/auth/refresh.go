package auth

import (
	"net/http"

	"dv/pkg/errors"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		errors.WriteJSONError(w, "missing refresh token", http.StatusUnauthorized)
		return
	}

	email, ok := refreshStore[cookie.Value]
	if !ok {
		errors.WriteJSONError(w, "invalid refresh token", http.StatusUnauthorized)
		return
	}

	accessToken, err := GenerateJWT(email)
	if err != nil {
		errors.WriteJSONError(w, "failed to generate new access token", http.StatusInternalServerError)
		return
	}

	// Optionally rotate refresh token
	newRefresh, err := generateRefreshToken()
	if err != nil {
		errors.WriteJSONError(w, "refresh rotation failed", http.StatusInternalServerError)
		return
	}

	delete(refreshStore, cookie.Value)
	refreshStore[newRefresh] = email

	setAuthCookies(w, accessToken, newRefresh)
	w.Write([]byte(`{"message":"refreshed successfully"}`))
}
