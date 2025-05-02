package repository

import (
	"database/sql"
	"real_time_forum/internal/model"
)

// UsersRepository يحدّد الدوال المتاحة للتعامل مع جدول المستخدمين.
type UsersRepository interface {
	RegisterUser(u *model.User) error
}

// UsersRepositorySqlite هو التنفيذ الخاص بقاعدة SQLite.
type UsersRepositorySqlite struct {
	DB *sql.DB
}

// RegisterUser يضيف مستخدماً جديداً إلى قاعدة البيانات.
func (r *UsersRepositorySqlite) RegisterUser(u *model.User) error {
	const query = `
		INSERT INTO users (nick_name, age, gender, first_name, last_name, email, password)
		VALUES (?, ?, ?, ?, ?, ?, ?);
	`
	_, err := r.DB.Exec(query,
		u.NickName, u.Age, u.Gender,
		u.FirstName, u.LastName, u.Email, u.Password,
	)
	return err
}
