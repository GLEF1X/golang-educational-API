package dto

import (
	"github.com/GLEF1X/golang-educational-API/domain/models"
	"time"
)

// User TODO: add advanced validation to models
type User struct {
	ID          int64      `json:"id,omitempty"`
	Name        string     `json:"full_name"`
	City        string     `json:"city"`
	ZipCode     string     `json:"zipcode"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Status      string     `json:"status"`
}

func DomainUserFromDTO(dtoUser *User) *models.User {
	return &models.User{
		Status:      dtoUser.Status,
		Name:        dtoUser.Name,
		City:        dtoUser.City,
		ZipCode:     dtoUser.ZipCode,
		DateOfBirth: dtoUser.DateOfBirth,
	}
}
