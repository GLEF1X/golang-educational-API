package interfaces

import "github.com/GLEF1X/golang-educational-API/domain/models"

type ICustomerRepository interface {
	FindAll() ([]models.User, error)
	AddOne(customer *models.User)
}
