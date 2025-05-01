package router

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	Routes map[string]map[string]http.HandlerFunc
}

// Instantiate a new router;
func NewRouter() *Router {
	return &Router{
		Routes: make(map[string]map[string]http.HandlerFunc),
	}
}

// Add a new route:
func (router *Router) AddRoute(method, path string, handler http.HandlerFunc) {
	if _, ok := router.Routes[method]; !ok {
		router.Routes[method] = make(map[string]http.HandlerFunc)
	}
	router.Routes[method][path] = handler
}

// make the router satisfy the http handler:
// Handle incoming HTTP requests
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := strings.ToLower(r.Method)
	path := r.URL.Path
	fmt.Printf("method: %v, path: %v", method, path)
	if handlers, ok := router.Routes[method]; ok {
		if handler, ok := handlers[path]; ok {
			handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
}
