package service

import (
	"movie-reservation-system/models"
)

type AuthService interface {
	Login(*models.UserLoginRequest) (*models.TokenDB, error)

	Logout(string) (*models.TokenDB, error)
}
