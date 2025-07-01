package auth

import (
	"net/http"

	"dv/pkg/errors"
)

func (auth *Auth) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		errors.WriteJSONError(w, "missing refresh token", http.StatusUnauthorized)
		return
	}

	email, ok := auth.RefreshStore[cookie.Value]
	if !ok {
		errors.WriteJSONError(w, "invalid refresh token", http.StatusUnauthorized)
		return
	}

	accessToken, err := auth.GenerateJWT(email)
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

	delete(auth.RefreshStore, cookie.Value)
	auth.RefreshStore[newRefresh] = email

	auth.SetAuthCookies(w, accessToken, newRefresh)
	w.Write([]byte(`{"message":"refreshed successfully"}`))
}
