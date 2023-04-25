package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("register"); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
		}
	}(db)

	repository := repositories.NewRepositoryUsers(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusCreated, user)

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
