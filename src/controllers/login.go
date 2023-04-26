package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

// Login authenticates a user in the API
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			return
		}
	}(db)

	repository := repositories.NewRepositoryUsers(db)
	userDB, err := repository.SearchByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VERIFY(user.Password, userDB.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userDB.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
