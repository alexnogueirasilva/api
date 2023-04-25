package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewRepositoryUsers(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	write, err := w.Write([]byte(fmt.Sprintf(
		"User created with ID %d",
		userID,
	)))
	if err != nil {
		return
	}

	fmt.Println(write)

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
