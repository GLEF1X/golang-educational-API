package models

import "time"

type User struct {
	ID          int64
	Name        string
	City        string
	ZipCode     string
	DateOfBirth *time.Time
	Status      string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
