package router

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	Routes map[string]http.HandlerFunc
}

// Instantiate a new router;
func NewRouter() *Router {
	return &Router{
		Routes: make(map[string]http.HandlerFunc),
	}
}

// Add a new route:
func (router *Router) AddRoute(method, path string, handler http.HandlerFunc) {
	key := fmt.Sprintf("%s%s", method, path)
	router.Routes[key] = handler
}

// make the router satisfy the http handler:
// Handle incoming HTTP requests
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := strings.ToLower(r.Method) +  r.URL.Path
	fmt.Printf("Routing ====> %v", params)
	if handler, ok := router.Routes[params]; ok {
		handler(w, r)
		return
	}

	http.NotFound(w, r)
}
