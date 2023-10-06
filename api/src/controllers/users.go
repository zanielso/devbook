package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	requestBody, error := io.ReadAll(r.Body)

	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error := user.FormatAndValidate("create"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersReposiotry(db)
	user.ID, error = repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func Update(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	requestBody, error := io.ReadAll(r.Body)

	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error := user.FormatAndValidate("update"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersReposiotry(db)
	error = repository.Update(userId, user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, error := strconv.ParseInt(parameters["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersReposiotry(db)
	user, error := repository.FindById(uint(userId))
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersReposiotry(db)
	users, error := repository.FindAll(nameOrNick)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
