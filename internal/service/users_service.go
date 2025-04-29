package service

import (
	"real_time_forum/internal/model"
	"real_time_forum/internal/repository"
)

// Create an interface to ease decoupling logic/implementation from functionality:
type UserServices interface {
	RegisterUser(age int, gender, firstName, lastName, email, passwrd string) error
}

// Create a structute to go ahead with the implementation:
type User_services struct {
	Repository repository.Users_repository
}

// register a new user:
func (user_serc User_services) RegisterUser(age int, gender, firstName, lastName, email, password string) error {
	user := &model.User{NickName: string(firstName[0]) + lastName, Age: age, FirstName: firstName, LastName: lastName, Email: email, Password: password}
	return user_serc.Repository.RegisterUser(user)
}
