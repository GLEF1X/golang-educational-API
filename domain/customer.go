package domain

import "time"

type Customer struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null;index"`
	City        string
	Zipcode     string
	DateOfBirth time.Time
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ICustomerRepository interface {
	FindAll() ([]Customer, error)
	AddOne(customer *Customer) error
}
