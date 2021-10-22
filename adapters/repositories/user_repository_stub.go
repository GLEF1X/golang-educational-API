package repositories

import (
	"github.com/GLEF1X/golang-educational-API/domain/models"
)

type UserRepositoryStub struct {
	customers []models.User
}

func (repo UserRepositoryStub) FindAll() ([]models.User, error) {
	return repo.customers, nil
}

func (repo UserRepositoryStub) AddOne(customer *models.User) {
	repo.customers = append(
		repo.customers,
		*customer,
	)
}

func NewCustomerRepositoryStub(initCustomers []models.User) *UserRepositoryStub {
	return &UserRepositoryStub{customers: initCustomers}
}
