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
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Println("[repositories.Create] Error closing statement:", err)
			return
		}
	}(statement)

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
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
			return
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
			return
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
			return
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
			return
		}
	}(statement)

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// SearchByEmail searches for a user by email in the database
func (repository Users) SearchByEmail(email string) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT id, password FROM devbook.users WHERE email = ?",
		email,
	)

	if err != nil {
		return models.User{}, err
	}

	defer func(lines *sql.Rows) {
		err := lines.Close()
		if err != nil {
			fmt.Println("[repositories.SearchByEmail] Error closing lines: ", err)
			return
		}
	}(lines)

	var user models.User

	if lines.Next() {
		if err := lines.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil

}

// Follow allows a user to follow another user
func (repository Users) Follow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"INSERT IGNORE INTO devbook.followers (user_id, follower_id) VALUES (?,?)",
	)
	if err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Println("[repositories.Follow] Error closing statement: ", err)
			return
		}
	}(statement)

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow allows a user to unfollow another user
func (repository Users) Unfollow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM devbook.followers WHERE user_id = ? AND follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Println("[repositories.Unfollow] Error closing statement: ", err)
			return
		}
	}(statement)

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// SearchFollowers searches for all followers of a user
func (repository Users) SearchFollowers(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		SELECT u.id, u.name, u.nickname, u.email, u.created_at FROM devbook.users u
		INNER JOIN devbook.followers f ON u.id = f.follower_id
		WHERE f.user_id = ?
	`, userID)

	if err != nil {
		return nil, err
	}

	defer func(lines *sql.Rows) {
		err := lines.Close()
		if err != nil {
			fmt.Println("[repositories.SearchFollowers] Error closing lines: ", err)
			return
		}
	}(lines)

	var followers []models.User

	for lines.Next() {
		var follower models.User

		if err := lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}
