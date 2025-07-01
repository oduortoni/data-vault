package controllers

import (
	"encoding/json"
	"net/http"

	"dv/internal/users"
	"dv/pkg/auth"
	"dv/pkg/errors"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(auth.UserKey).(*users.UserDTO)
	if !ok {
		errors.WriteJSONError(w, "Internal Server Error: could not retrieve user from context", http.StatusInternalServerError)
		return
	}

	data := map[string]string {
		"username": user.Username,
		"email":    user.Email,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		errors.WriteJSONError(w, "Internal Server Error: could not encode user data", http.StatusInternalServerError)
		return
	}
}
