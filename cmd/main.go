package main

import (
	"fmt"
	"log"
	"net/http"
	"real_time_forum/database"
	"real_time_forum/internal/handler"
	"real_time_forum/internal/repository"
	"real_time_forum/internal/router"
	"real_time_forum/internal/service"
)

func main() {
	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	
	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	
	// Initialize repositories
	userRepo := &repository.UsersRepo{DB: db}
	
	// Initialize services
	userService := service.NewUserService(userRepo)
	
	// Initialize handlers
	userHandler := &handler.UserHandler{
		UserService: userService,
	}
	
	// Initialize router
	r := router.NewRouter(userHandler)
	
	// Start server
	port := ":8080"
	fmt.Printf("Server started on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}