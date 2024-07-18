package model

type JobInfo struct {
	Model
	Company  string
	Checkin  string
	Checkout string
	Status   int
}
