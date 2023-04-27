package models

import (
	"errors"
	"strings"
)

// Publication struct for publications
type Publication struct {
	ID         uint64 `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorID   uint64 `json:"author_id,omitempty"`
	AuthorNick string `json:"author_nick,omitempty"`
	Likes      uint64 `json:"likes"`
	CreatedAt  string `json:"created_at,omitempty"`
}

// Prepare prepares the publication
func (publication *Publication) Prepare() error {
	if err := publication.validate(); err != nil {
		return err
	}

	publication.format()
	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("the title is required")
	}

	if publication.Content == "" {
		return errors.New("the content is required")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
