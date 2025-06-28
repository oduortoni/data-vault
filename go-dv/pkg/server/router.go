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
	if handler, ok := r.routes[req.URL.Path]; ok && req.Method == http.MethodGet {
		handler(w, req)
		return
	}

	http.NotFound(w, req)
}
