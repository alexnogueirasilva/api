package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Search searches for all users that have the name or nickname
func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repository.db.Query(
		"SELECT id, name, nickname, email, created_at FROM devbook.users WHERE name LIKE ? OR nickname LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer func(lines *sql.Rows) {
		err := lines.Close()
		if err != nil {
			fmt.Println("[repositories.Search] Error closing lines: ", err)
		}
	}(lines)

	var users []models.User

	for lines.Next() {
		var user models.User

		if err := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// SearchByID searches for a user by ID in the database
func (repository Users) SearchByID(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT id, name, nickname, email, created_at FROM devbook.users WHERE id = ?",
		ID,
	)

	if err != nil {
		return models.User{}, err
	}

	defer func(lines *sql.Rows) {
		err := lines.Close()
		if err != nil {
			fmt.Println("[repositories.SearchByID] Error closing lines: ", err)
		}
	}(lines)

	var userFound models.User

	if lines.Next() {
		if err := lines.Scan(
			&userFound.ID,
			&userFound.Name,
			&userFound.Nick,
			&userFound.Email,
			&userFound.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return userFound, nil

}

// Update updates a user in the database
func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE devbook.users SET name = ?, nickname = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Println("[repositories.Update] Error closing statement: ", err)
		}
	}(statement)

	if _, err := statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// Delete deletes a user in the database
func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM devbook.users WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Println("[repositories.Delete] Error closing statement: ", err)
		}
	}(statement)

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
