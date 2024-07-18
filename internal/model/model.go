package model

import "time"

type Model struct {
	ID        uint `gorm:"primaryKey;type:int;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
