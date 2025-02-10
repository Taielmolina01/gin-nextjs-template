package repository

import (
	"movie-reservation-system/models"
)

type AuthRepository interface {
	CreateToken(*models.TokenDB) (*models.TokenDB, error)

	GetToken(string) (*models.TokenDB, error)

	DeleteToken(*models.TokenDB) (*models.TokenDB, error)
}
