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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. send success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}
