package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
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

	if err = publication.Prepare(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

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
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

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
	publications, err := repository.GetPublications(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)
}

// GetPublication returns a publication
func GetPublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

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
	publication, err := repository.GetPublicationByID(publicationID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publication)
}

// UpdatePublication updates a publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

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
	publicationSaved, err := repository.GetPublicationByID(publicationID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if publicationSaved.AuthorID != userID {
		response.Error(w, http.StatusForbidden, errors.New("you cannot update a publication that is not yours"))
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

	if err = publication.Prepare(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(publicationID, publication); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// DeletePublication deletes a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

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
	publicationSaved, err := repository.GetPublicationByID(publicationID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if publicationSaved.AuthorID != userID {
		response.Error(w, http.StatusForbidden, errors.New("you cannot delete a publication that is not yours"))
		return
	}

	if err = repository.Delete(publicationID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// GetPublicationsByUser returns all publications for a user
func GetPublicationsByUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

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
	publications, err := repository.GetPublicationsByUser(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)
}

// LikePublication likes a publication
func LikePublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

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
	if err = repository.Like(publicationID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
