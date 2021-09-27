package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// create a new endpoint route
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/tinyurl", generateTinyUrl).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Tiny URL: Entry Point")
	handleRequests()
}
