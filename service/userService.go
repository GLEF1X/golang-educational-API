package service

import (
	"github.com/GLEF1X/golang-educational-API/domain"
	"log"
)

type CustomerService struct {
	repository domain.ICustomerRepository
}

func (s *CustomerService) GetAllCustomers() []domain.Customer {
	customers, err := s.repository.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	return customers
}

func (s *CustomerService) AddCustomer(customer *domain.Customer) error {
	return s.repository.AddOne(customer)
}

func NewCustomerService(repository domain.ICustomerRepository) *CustomerService {
	return &CustomerService{repository: repository}
}
