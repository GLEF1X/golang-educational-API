package repositories

import "github.com/GLEF1X/golang-educational-API/domain"

type CustomerRepositoryStub struct {
	customers []domain.Customer
}

func (repo CustomerRepositoryStub) FindAll() ([]domain.Customer, error) {
	return repo.customers, nil
}

func (repo CustomerRepositoryStub) AddOne(customer *domain.Customer) error {
	repo.customers = append(repo.customers, *customer)
	return nil
}

func NewCustomerRepositoryStub(initCustomers []domain.Customer) *CustomerRepositoryStub {
	return &CustomerRepositoryStub{customers: initCustomers}
}
