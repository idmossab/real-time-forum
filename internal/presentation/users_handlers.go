package presentation

import (
	"net/http"
	"real_time_forum/internal/service"
)

// Create a structure for the task haandler:
type UsersHandler struct {
	Service service.UserServices
}

// Handle the user registration:
func (handler UsersHandler) UserRegistrationHandler(rs http.ResponseWriter, rq *http.Request){
	
}
