package model

import "time"

type User struct {
	Model
	FirstName     string
	LastName      string
	DateOfBirth   *time.Time
	Email         string
	ContactNumber string
	Notes         string
	Status        int
}
