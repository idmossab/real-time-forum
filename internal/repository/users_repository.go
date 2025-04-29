package repository

import (
	"database/sql"
	"real_time_forum/internal/model"
)

// Create an interface for decoupled implementation:
type UsersRepository interface {
	RegisterUser(user *model.User) error
}

// Create the the structure that will implement the userRepository interface:
type Users_repository struct {
	Database *sql.DB
}

// start the struct/interface implementation:
func (user_repo Users_repository) RegisterUser(user *model.User)error{
	
return nil
}