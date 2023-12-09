package bootstrap

import (
	"database/sql"
	"fmt"
)

// ConnectDB establishes a connection to the SQLite database
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// DB is a global variable that holds the connection to the database
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}

	DB = db

	return nil
}

// CloseDB closes the connection to the database
func CloseDB() {
	if DB != nil {
		defer DB.Close()
	}
}
