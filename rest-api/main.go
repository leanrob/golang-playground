package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// main function to boot up everything
func main() {
	router := mux.NewRouter()

	// Server the static index page on the default page load
	router.Handle("/", http.FileServer(http.Dir("./views/")))

	// Server all static resources that are in the static folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Create 3 people for start of API
	people = append(people, Person{ID: "1", Firstname: "Rob", Lastname: "B", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Kelly", Lastname: "G", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Jim", Lastname: "T", Address: &Address{City: "City Z", State: "State Y"}})

	// Create API routes
	router.HandleFunc("/people/all", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}
