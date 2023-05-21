package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
}

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reading all users..."))
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