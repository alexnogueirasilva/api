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
	"strings"
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
		return
	}

	response.JSON(w, http.StatusCreated, user)

}

// SearchUsers searches for all users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nicknameOrName := strings.ToLower(r.URL.Query().Get("user"))
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

	repository := repositories.NewRepositoryUsers(db)
	users, err := repository.Search(nicknameOrName)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)

}

// SearchUser searches for a user
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
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

	repository := repositories.NewRepositoryUsers(db)
	user, err := repository.SearchByID(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, user)

}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDInToken {
		response.Error(w, http.StatusForbidden, errors.New("you cannot update a user that is not yours"))
		return
	}

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

	if err = user.Prepare("update"); err != nil {
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

	repository := repositories.NewRepositoryUsers(db)
	if err = repository.Update(userID, user); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDInToken {
		response.Error(w, http.StatusForbidden, errors.New("you cannot delete a user that is not yours"))
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

	repository := repositories.NewRepositoryUsers(db)
	if err = repository.Delete(userID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// FollowUser follows a user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		response.Error(w, http.StatusForbidden, errors.New("you cannot follow yourself"))
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

	repository := repositories.NewRepositoryUsers(db)
	if err = repository.Follow(userID, followerID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser unfollows a user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		response.Error(w, http.StatusForbidden, errors.New("you cannot unfollow yourself"))
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

	repository := repositories.NewRepositoryUsers(db)
	if err = repository.Unfollow(userID, followerID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)
}
