package database

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Connect to my database:
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "real_time_forum.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, nil
}