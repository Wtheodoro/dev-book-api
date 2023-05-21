package router

import "github.com/gorilla/mux"

// function  Generate returns a router with configured routes
func Generate() *mux.Router {
	return mux.NewRouter()
}