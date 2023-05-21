package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando api")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5050", r))
}