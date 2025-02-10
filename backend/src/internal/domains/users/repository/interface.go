package repository

import (
	"movie-reservation-system/models"
)

type UserRepository interface {
	CreateUser(*models.UserDB) (*models.UserDB, error)

	GetUser(string) (*models.UserDB, error)

	UpdateUser(*models.UserDB) (*models.UserDB, error)

	DeleteUser(*models.UserDB) (*models.UserDB, error)
}
