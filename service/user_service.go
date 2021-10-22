package service

import (
	"github.com/GLEF1X/golang-educational-API/core/logging"
	"github.com/GLEF1X/golang-educational-API/domain/interfaces"
	"github.com/GLEF1X/golang-educational-API/domain/models"
	"github.com/GLEF1X/golang-educational-API/dto"
)

type UserService struct {
	repository interfaces.ICustomerRepository
}

func (s *UserService) GetAllUsers() []models.User {
	users, err := s.repository.FindAll()
	if err != nil {
		logging.GetLogger().Error(err)
	}
	return users
}

func (s *UserService) AddUser(user *dto.User) {
	s.repository.AddOne(dto.DomainUserFromDTO(user))
}

func NewCustomerService(repository interfaces.ICustomerRepository) *UserService {
	return &UserService{repository: repository}
}
