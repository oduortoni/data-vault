package auth

import (
	"context"
	"net/http"

	"dv/pkg/errors"
)

type contextKey string

// UserKey is the key used to store the user object in the request context.
const UserKey = contextKey("user")

// AuthMiddleware protects routes by validating the JWT in the cookie.
// If the token is valid, it fetches the user data and stores it in the request context.
func (auth *Auth) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(auth.TokenName)
		if err != nil {
			// If no cookie, treat as unauthorized
			errors.WriteJSONError(w, "unauthorized: missing auth token", http.StatusUnauthorized)
			return
		}

		// Parse the JWT to get the user's email
		email, err := auth.ParseJWT(cookie.Value)
		if err != nil {
			errors.WriteJSONError(w, "unauthorized: invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Fetch the full user details from the database using the email from the token.
		user, found := auth.UserService.Exists(email)
		if !found {
			// This can happen if the user was deleted after the token was issued.
			errors.WriteJSONError(w, "unauthorized: user not found", http.StatusUnauthorized)
			return
		}

		// Store the user object in the request context and call the next handler.
		ctx := context.WithValue(r.Context(), UserKey, user)
		next(w, r.WithContext(ctx))
	}
}
