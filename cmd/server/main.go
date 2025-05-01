package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"real_time_forum/database"
	"real_time_forum/internal/presentation"
	"real_time_forum/internal/repository"
	"real_time_forum/internal/router"
	"real_time_forum/internal/service"
)

var databaseConnection *sql.DB
var mainError error


func init() {
	databaseConnection, mainError = database.Connect()
	if mainError == nil {
		mainError = database.Migrate(databaseConnection)
	}
}

func main() {
	if mainError != nil {
		fmt.Println("Error connecting to database: %w", mainError)
		return
	}
	defer databaseConnection.Close()
	fmt.Println("connected successfully")

	userRepository := repository.Users_repository{Database: databaseConnection}
	userService := service.User_services{Repository: userRepository}
	userHandler := presentation.UsersHandler{Service: userService}
	
	mainRouter := router.NewRouter()
	mainRouter.AddRoute("post", "/add", userHandler.UserRegistrationHandler)
	fmt.Println(mainRouter.Routes)
	fmt.Println("listenning on port: http://localhost:8080/")
	mainError = http.ListenAndServe(":8080", mainRouter)
	if mainError != nil {
		fmt.Println("Error: %w", mainError)
		return
	}
	}

