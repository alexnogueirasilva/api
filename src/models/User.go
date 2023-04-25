package models

import (
	"errors"
	"strings"
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare calls the functions to validate and format the user received
func (user *User) Prepare(step string) error {
	if err := user.validate(); err != nil {
		return err
	}

	if err := user.format(); err != nil {
		return err
	}

	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("the name is required")
	}

	if user.Nick == "" {
		return errors.New("the nick is required")
	}
	if user.Email == "" {
		return errors.New("the email is required")
	}

	if user.Password == "" {
		return errors.New("the password is required")
	}

	return nil
}

func (user *User) format() error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	return nil
}
