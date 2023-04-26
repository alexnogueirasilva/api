package models

import (
	"api/src/security"
	"errors"
	"github.com/badoux/checkmail"
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
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("the name is required")
	}
	if user.Nick == "" {
		return errors.New("the nick is required")
	}
	if user.Email == "" {
		return errors.New("the email is required")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("the email is invalid")
	}
	if step == "register" && user.Password == "" {
		return errors.New("the password is required")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hash, err := security.HASH(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hash)
	}

	return nil
}
