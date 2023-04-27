package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publications RepositoryPublications is a repository for publications
type Publications struct {
	db *sql.DB
}

func NewRepositoryPublications(db *sql.DB) *Publications {
	return &Publications{db}
}

// Create creates a new publication
func (repository Publications) Create(publications models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO devbook.publications (title, content, author_id) VALUES (?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			return
		}
	}(statement)

	result, err := statement.Exec(publications.Title, publications.Content, publications.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}
