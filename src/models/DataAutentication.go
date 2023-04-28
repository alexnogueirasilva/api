package models

// Authentication stores the token and the user ID
type Authentication struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
