package model

import "time"

type Model struct {
	ID        uint `gorm:"primaryKey;type:int;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ListQuery struct {
	Limit  int
	Offset int
	And    []AndQuery
	Sorts  []SortQuery
}

type AndQuery struct {
	Or []OrQuery
}

type OrQuery struct {
	Cond  string
	Value any
}

type SortQuery struct {
	Field string
	Dir   string
}
