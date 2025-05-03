package repository

import (
	"database/sql"
	"fmt"
	"real_time_forum/internal/model"
)

type UsersRepository interface {
	RegisterUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
}

type UsersRepo struct {
	DB *sql.DB
}

func (ur *UsersRepo) RegisterUser(user *model.User) error {
	query := `INSERT INTO users(nick_name, age, gender, first_name, last_name, email, password)
			  VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := ur.DB.Exec(query, user.NickName, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UsersRepo) GetUserByEmail(email string) (*model.User, error) {
	query := `SELECT id, nick_name, age, gender, first_name, last_name, email, password 
			  FROM users WHERE email = ?`

	var user model.User
	err := ur.DB.QueryRow(query, email).Scan(
		&user.Id, &user.NickName, &user.Age, &user.Gender,
		&user.FirstName, &user.LastName, &user.Email, &user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("error querying user by email: %w", err)
	}

	return &user, nil
}
