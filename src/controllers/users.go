package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	database, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)
	user.ID, error = repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	database, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer database.Close()

	repository := repositories.NewUsersRepository(database)
	users, error := repository.Search(nameOrNick)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("reading one user..."))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating user..."))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting user..."))
}