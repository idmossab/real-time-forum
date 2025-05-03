package handler

import (
	"encoding/json"
	"net/http"
	"real_time_forum/internal/model"
	"real_time_forum/internal/service"
)

type UserHandler struct {
	UserService service.UserServices
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. decode JSON
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 2. call service
	err = h.UserService.RegisterUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	response := map[string]string{
		"status":  "success",
		"message": "User registered successfully",
	}
	
	json.NewEncoder(w).Encode(response)
}

// Login handles user authentication
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse login credentials
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Authenticate user
	user, err := h.UserService.AuthenticateUser(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	
	response := map[string]interface{}{
		"status":  "success",
		"message": "Login successful",
		"user": map[string]interface{}{
			"id":        user.Id,
			"nick_name": user.NickName,
			"email":     user.Email,
		},
	}
	
	json.NewEncoder(w).Encode(response)
}