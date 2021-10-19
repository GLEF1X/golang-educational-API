package repositories

import (
	"github.com/GLEF1X/golang-educational-API/domain"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	connection *gorm.DB
}

func (repository *CustomerRepository) FindAll() ([]domain.Customer, error) {
	var records []domain.Customer
	_, err := repository.connection.Find(&records).Rows()
	return records, err
}

func (repository *CustomerRepository) AddOne(customer *domain.Customer) error {
	result := repository.connection.Create(customer)
	return result.Error
}

func NewCustomerRepository(conn *gorm.DB) *CustomerRepository {
	return &CustomerRepository{connection: conn}
}
