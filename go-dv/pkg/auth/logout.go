package auth

import "net/http"

func Logout(w http.ResponseWriter, r *http.Request) {
	// Clear cookies
	http.SetCookie(w, &http.Cookie{Name: "access_token", Value: "", Path: "/", MaxAge: -1})
	http.SetCookie(w, &http.Cookie{Name: "refresh_token", Value: "", Path: "/auth/refresh", MaxAge: -1})

	// Invalidate refresh token
	if cookie, err := r.Cookie("refresh_token"); err == nil {
		delete(refreshStore, cookie.Value)
	}

	w.Write([]byte(`{"message":"logged out"}`))
}
