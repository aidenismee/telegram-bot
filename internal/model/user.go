package model

type User struct {
	Model
	FirstName     string
	LastName      string
	DateOfBirth   string
	Email         string
	ContactNumber string
	Notes         string
	Status        int
}
