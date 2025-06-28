package server

import (
	"net/http"
)

type Router struct {
	routes map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) register(pattern string, handler http.HandlerFunc) {
	r.routes[pattern] = handler
}

// implement the ServeHTTP so the Router satisfies http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for pattern, handler := range r.routes {
		if req.Method == http.MethodGet && pathMatch(pattern, req.URL.Path) {
			handler(w, req)
			return
		}
	}

	http.NotFound(w, req)
}

// pathMatch checks for exact or prefix matches
func pathMatch(pattern, path string) bool {
	// Exact match
	if pattern == path {
		return true
	}

	// Prefix match for paths like "/static/"
	if len(pattern) > 1 && pattern[len(pattern)-1] == '/' && len(path) >= len(pattern) && path[:len(pattern)] == pattern {
		return true
	}

	return false
}
