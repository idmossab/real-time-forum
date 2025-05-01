package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real_time_forum/internal/model"
	"real_time_forum/internal/service"
)

// Create a structure for the task haandler:
type UsersHandler struct {
	Service service.UserServices
}

// Handle the user registration:
func (handler UsersHandler) UserRegistrationHandler(wr http.ResponseWriter, rq *http.Request) {
	if rq.Method == "POST" {
		user := model.User{}
		if err := json.NewDecoder(rq.Body).Decode(&user); err != nil {
			http.Error(wr, "Invalid request body", http.StatusBadRequest)
			return
		}
		fmt.Println(user)
		wr.Header().Set("Content-Type", "application/json")
		json.NewEncoder(wr).Encode(&user)
		return
	}
	http.Error(wr, "Registration has failed.", http.StatusInternalServerError)
}
