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
func NewRouter()*Router{
return &Router{
	Routes: make(map[string]map[string]http.HandlerFunc),
}
}

// Add a new route:
func(router *Router) AddRoute(method, path string, handler http.HandlerFunc){
router.Routes[method] = make(map[string]http.HandlerFunc)
router.Routes[method][path] = handler
}

// make the router satisfy the http handler:
func (router *Router)ServeHTTP(wr http.ResponseWriter, rq *http.Request){
	fmt.Printf("the path: %s => the method: %s", rq.Method, rq.URL.Path[1:])
if router.Routes[strings.ToLower(rq.Method)] == nil {
	http.Error(wr, "Method not allowed", http.StatusMethodNotAllowed)
	return
} else if router.Routes[rq.Method][rq.URL.Path[1:]] == nil {
	http.NotFound(wr, rq)
	return
}else {
	router.Routes[rq.Method][rq.URL.Path](wr, rq)
}
}