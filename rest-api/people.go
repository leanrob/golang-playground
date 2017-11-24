package main

import (

)

// The person Type (more like an object)
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// Create 3 people for start of API
var people = []Person{
	Person{ID: "1", Firstname: "Rob", Lastname: "B", Address: &Address{City: "City X", State: "State X"}},
	Person{ID: "2", Firstname: "Kelly", Lastname: "G", Address: &Address{City: "City Z", State: "State Y"}},
	Person{ID: "3", Firstname: "Jim", Lastname: "T", Address: &Address{City: "City Z", State: "State Y"}},
}

