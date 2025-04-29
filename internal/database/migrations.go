package database

import (
	"database/sql"
	"fmt"
)

const schema = `CREATE TABLE IF NOT EXISTS Users(
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    NICK_NAME TEXT UNIQUE NOT NULL,
    AGE INTEGER CHECK(AGE > 0 AND AGE < 100),
    GENDER TEXT NOT NULL, 
    FIRST_NAME TEXT NOT NULL,
    LAST_NAME TEXT NOT NULL,
    E_MAIL TEXT NOT NULL UNIQUE,
    PASSWORD TEXT NOT NULL
);`

func Migrate(db *sql.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("error migrating: %w", err)
	}
	return nil
}
