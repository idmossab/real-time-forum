package service

import (
	"errors"
	"real_time_forum/internal/model"
	"real_time_forum/internal/repository"
	"real_time_forum/internal/utils"

)

type UserServices interface {
	RegisterUser(user *model.User) error
}

type UserService struct {
	Repository repository.UsersRepository
}

func (us *UserService) RegisterUser(user *model.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}
		// Hash the password before storing
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	return us.Repository.RegisterUser(user)
}

func NewUserService(repo repository.UsersRepository) *UserService {
	return &UserService{
		Repository: repo,
	}
}
