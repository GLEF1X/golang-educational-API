package repositories

import (
	"context"
	"github.com/GLEF1X/golang-educational-API/core/logging"
	"github.com/GLEF1X/golang-educational-API/domain/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func (repository *UserRepository) FindAll() ([]models.User, error) {
	var records []models.User
	err := pgxscan.Select(context.Background(), repository.db, &records, "SELECT * FROM users")
	if err != nil {
		logging.GetLogger().Error(err)
	}
	return records, nil
}

func (repository *UserRepository) AddOne(user *models.User) {
	_, err := repository.db.Exec(
		context.Background(),
		`INSERT INTO users(name, city, zip_code, date_of_birth, status)
			 VALUES ($1, $2, $3, $4, $5)`,
		user.Name, user.City, user.ZipCode, user.DateOfBirth, user.Status,
	)
	if err != nil {
		logging.GetLogger().Error(err)
	}
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}
