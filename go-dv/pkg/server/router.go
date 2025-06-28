package server

import (
	"net/http"
)
type routeKey struct {
	Method string
	Path   string
}

type Router struct {
	routes map[routeKey]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[routeKey]http.HandlerFunc),
	}
}

func (r *Router) register(method string, pattern string, handler http.HandlerFunc) {
	r.routes[routeKey{Method: method, Path: pattern}] = handler
}
// implement the ServeHTTP so the Router satisfies http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for key, handler := range r.routes {
		if key.Method == req.Method && pathMatch(key.Path, req.URL.Path) {
			handler(w, req)
			return
		}
	}
	http.Redirect(w, req, "/", http.StatusFound)
}


// pathMatch checks for exact or prefix matches
func pathMatch(pattern, path string) bool {
	// exact match
	if pattern == path {
		return true
	}

	// prefix match for paths like "/static/"
	if len(pattern) > 1 && pattern[len(pattern)-1] == '/' && len(path) >= len(pattern) && path[:len(pattern)] == pattern {
		return true
	}

	return false
}
