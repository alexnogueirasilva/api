package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

// CreatePublication creates a new publication
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(bodyRequest, &publication); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = userID

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
		}
	}(db)

	repository := repositories.NewRepositoryPublications(db)
	publication.ID, err = repository.Create(publication)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, publication)
}

// GetPublications returns all publications
func GetPublications(w http.ResponseWriter, r *http.Request) {

}

// GetPublication returns a publication
func GetPublication(w http.ResponseWriter, r *http.Request) {

}

// UpdatePublication updates a publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

// DeletePublication deletes a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
