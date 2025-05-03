package router

import (
	"net/http"
	"real_time_forum/internal/handler"
)

// NewRouter creates and configures a new router
func NewRouter(userHandler *handler.UserHandler) http.Handler {
	// Create new ServeMux
	mux := http.NewServeMux()
	
	// Register API routes
	mux.HandleFunc("/api/register", userHandler.Register)
	
	// You can add more routes here as your application grows
	
	return mux
}