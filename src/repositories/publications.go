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

// GetPublicationByID returns a publication by ID
func (repository Publications) GetPublicationByID(publicationID uint64) (models.Publication, error) {
	line, err := repository.db.Query(`
		SELECT p.*, u.nickname FROM devbook.publications p
		INNER JOIN devbook.users u ON u.id = p.author_id
		WHERE p.id = ?
	`, publicationID)
	if err != nil {
		return models.Publication{}, err
	}
	defer func(line *sql.Rows) {
		err := line.Close()
		if err != nil {
			return
		}
	}(line)

	var publication models.Publication
	if line.Next() {
		if err = line.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}

// GetPublications returns all publications
func (repository Publications) GetPublications(userID uint64) ([]models.Publication, error) {
	lines, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.nickname FROM devbook.publications p
		INNER JOIN devbook.users u ON u.id = p.author_id
		INNER JOIN devbook.followers f ON p.author_id = f.user_id
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 DESC
	`, userID, userID,
	)

	if err != nil {
		return nil, err
	}
	defer func(lines *sql.Rows) {
		err := lines.Close()
		if err != nil {
			return
		}
	}(lines)

	var publications []models.Publication
	for lines.Next() {
		var publication models.Publication
		if err = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Update updates a publication
func (repository Publications) Update(publicationID uint64, publication models.Publication) error {
	statement, err := repository.db.Prepare(
		"UPDATE devbook.publications SET title = ?, content = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			return
		}
	}(statement)

	if _, err = statement.Exec(publication.Title, publication.Content, publicationID); err != nil {
		return err
	}

	return nil
}
