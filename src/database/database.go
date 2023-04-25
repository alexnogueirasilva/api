package database

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connect opens a connection with the database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return db, nil
}
