package auth

import (
	"context"
	"dv/pkg/errors"
	"net/http"
)

type contextKey string

const UserKey = contextKey("user_email")

// AuthMiddleware protects routes by validating the JWT in the cookie
func (auth *Auth) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(auth.TokenName)
		if err != nil {
			errors.WriteJSONError(w, "unauthorized: missing auth token", http.StatusUnauthorized)
			return
		}

		email, err := auth.ParseJWT(cookie.Value)
		if err != nil {
			errors.WriteJSONError(w, "unauthorized: invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, email)
		next(w, r.WithContext(ctx))
	}
}
