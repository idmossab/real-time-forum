package router

import (
	"net/http"
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
if router.Routes[rq.Method] == nil {
	http.Error(wr, "Method not allowed", http.StatusMethodNotAllowed)
	return
} else if router.Routes[rq.Method][rq.URL.Path] == nil {
	http.NotFound(wr, rq)
	return
}else {
	router.Routes[rq.Method][rq.URL.Path](wr, rq)
}
}