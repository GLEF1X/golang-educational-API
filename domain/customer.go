package domain

import "time"

type Customer struct {
	Id          int
	Name        string
	City        string
	Zipcode     string
	DateOfBirth time.Time
	Status      string
}

type ICustomerRepository interface {
	FindAll() ([]Customer, error)
}
