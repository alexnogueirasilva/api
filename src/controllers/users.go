package controllers

import "net/http"

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user!"))
}

// SearchUsers searches for all users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching users!"))
}

// SearchUser searches for a user
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching user!"))
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user!"))
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user!"))
}
