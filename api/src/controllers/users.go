package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}

	repository := repositories.NewUsersReposiotry(db)
	userId, error := repository.Create(user)
	if error != nil {
		log.Fatal(error)
	}

	w.Write([]byte(fmt.Sprintf("Id Inserted: %d", userId)))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user"))
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All users"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
