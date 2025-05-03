package service

import (
	"errors"
	"real_time_forum/internal/model"
	"real_time_forum/internal/repository"
	"real_time_forum/internal/utils"
)

type UserServices interface {
	RegisterUser(user *model.User) error
	AuthenticateUser(email, password string) (*model.User, error)
}

type UserService struct {
	Repository repository.UsersRepository
}

func NewUserService(repo repository.UsersRepository) *UserService {
	return &UserService{
		Repository: repo,
	}
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

// AuthenticateUser verifies user credentials and returns the user if valid
func (us *UserService) AuthenticateUser(email, password string) (*model.User, error) {
	// Input validation
	if email == "" {
		return nil, errors.New("email is required")
	}
	if password == "" {
		return nil, errors.New("password is required")
	}

	// Get user by email
	user, err := us.Repository.GetUserByEmail(email)
	if err != nil {
		// Log the error but don't expose details to client
		// fmt.Printf("Login error: %v\n", err)
		return nil, errors.New("invalid email or password")
	}

	// Check if password matches
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}