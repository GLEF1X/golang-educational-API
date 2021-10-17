package service

import (
	"fasthttp_restful/domain"
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

func NewCustomerService(repository domain.ICustomerRepository) *CustomerService {
	return &CustomerService{repository: repository}
}
