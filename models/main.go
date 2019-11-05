package models

import "time"

// User is defined model of User entities
type User struct {
	Name string
	Email string
	Password string
	PhoneNumber string
	Location Location
}

// Lawyer is defined model of Lawyer entities
type Lawyer struct {
	User User
	FullName    string
	Language string
	Experiences int
	Rate        float64
}

// Location is defined model of Location entities
type Location struct {
	City, Province, Country string
}

// Transaction is defined model of Transaction entities
type Transaction struct {
	Length float64
	Time time.Time
	User User
	Lawyer Lawyer
}

// Review is defined model of Review entities
type Review struct {
	User User
	Lawyer Lawyer
	Transaction Transaction
	Rate int
	Title string
	Description string
	Time time.Time
}