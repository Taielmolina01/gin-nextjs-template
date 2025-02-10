package service

import (
	"movie-reservation-system/models"
)

type UserService interface {
	CreateUser(*models.UserRequest) (*models.UserDB, error)

	GetUser(string) (*models.UserDB, error)

	UpdateUser(string, *models.UserUpdateRequest) (*models.UserDB, error)

	UpdateUserPassword(string, *models.UserUpdatePasswordRequest) (*models.UserDB, error)

	DeleteUser(string) (*models.UserDB, error)
}
