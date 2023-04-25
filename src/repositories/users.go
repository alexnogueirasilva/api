package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represents a repositories of users
type Users struct {
	db *sql.DB
}

// NewRepositoryUsers creates a new repositories that implements the users.Repository interface
func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

// Create inserts a user in the database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"	INSERT INTO devbook.users (name, nickname, email, password) VALUES (?,?,?,?)",
	)
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, nil
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastInsertedID), nil
}
