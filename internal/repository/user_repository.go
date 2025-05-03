package repository

import (
	"database/sql"
	"real_time_forum/internal/model"
)

type UsersRepository interface {
	RegisterUser(user *model.User) error
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
