package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
)

// main function to boot up everything
func main() {
	router := mux.NewRouter()

	// Server the static index page on the default page load
	router.Handle("/", http.FileServer(http.Dir("./views/")))

	// Server all static resources that are in the static folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Create 3 people for start of API
	people = append(people, Person{ID: "4", Firstname: "This", Lastname: "Guy", Address: &Address{City: "City Z", State: "State Y"}})

	// Create API routes
	router.HandleFunc("/people/all", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	// Middleware to wrap our router in a log so the logger is called first every time
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
