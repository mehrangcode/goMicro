package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", os.Getenv("dbName"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func InitDB() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	DB = db
	MigrateDB()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func CloseDB() {
	if DB != nil {
		defer DB.Close()
	}
}

func MigrateDB() error {
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at CURRENT_TIMESTAMP NOT NULL DEFAULT (now())
		);
		`
	_, err := DB.Exec(query)
	errorHandler(err)
	query = `
		CREATE TABLE IF NOT EXISTS courses(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			teacherId TEXT NOT NULL,
			created_at CURRENT_TIMESTAMP NOT NULL DEFAULT (now()),
			FOREIGN KEY (teacherId) 
      		REFERENCES users (teacherId) 
         	ON DELETE CASCADE 
         	ON UPDATE NO ACTION
		);
		`
	_, err = DB.Exec(query)
	errorHandler(err)
	// query = `
	// ALTER TABLE users ADD COLUMN IF NOT EXISTS created_at TEXT;
	// 	`
	// _, err = DB.Exec(query)
	// errorHandler(err)
	// query = `
	// ALTER TABLE courses ADD COLUMN IF NOT EXISTS created_at TEXT;
	// 	`
	// _, err = DB.Exec(query)
	// errorHandler(err)
	return err
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
