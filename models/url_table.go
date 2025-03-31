package models

import (
	"database/sql"
)

func CreateURLTable(db *sql.DB) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS urls (
		id INT AUTO_INCREMENT PRIMARY KEY,
		original_url TEXT NOT NULL,
		short_url VARCHAR(10) NOT NULL UNIQUE
	)`
	_, err := db.Exec(createTableQuery)

	return err
}
