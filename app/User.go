package main

import "time"

// User is the structure representation of a User entry in the database
// teh value member since is defined in the creation of the user, cannot be changed
type User struct {
	ID           int       `json:"id"`
	LastName     string    `json:"last_name"`
	FirstName    string    `json:"first_name"`
	BirthDate    time.Time `json:"birth_date"`
	Mail         string    `json:"mail"`
	Phone        string    `json:"phone"`
	memeberSince time.Time
	Biography    string `json:"biography"`
	PhotoURL     string `json:"photo_url"`
}
