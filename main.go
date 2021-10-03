package main

import (
	"create-tiny-url/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	// create a new instance of a mux router
	r := mux.NewRouter().StrictSlash(true)

	// create a new endpoint route
	r.HandleFunc("/", services.HomePage)
	r.HandleFunc("/tinyurl", services.GetTinyUrl).Methods("POST")
	r.HandleFunc("/normalurl", services.GenerateNormalUrl).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("Tiny URL: Entry Point")
	handleRequests()
}
