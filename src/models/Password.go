package models

// Password represents a password
type Password struct {
	New    string `json:"new" validate:"required,min=6"`
	Actual string `json:"actual" validate:"required,min=6"`
}
