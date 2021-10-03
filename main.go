package main

import (
	"create-tiny-url/internal"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	// create a new instance of a mux router
	r := mux.NewRouter().StrictSlash(true)

	// create a new endpoint route
	r.HandleFunc("/", internal.HomePage)
	r.HandleFunc("/tinyurl", internal.GetTinyUrl).Methods("POST")
	r.HandleFunc("/normalurl", internal.GenerateNormalUrl).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("Tiny URL: Entry Point")
	handleRequests()
}
