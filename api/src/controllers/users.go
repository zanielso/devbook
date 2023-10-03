package controllers

import "net/http"

func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
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
