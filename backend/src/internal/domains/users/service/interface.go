package service

import (
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
)

type UserService interface {
	CreateUser(*models.UserRequest) (*models.UserCRUDResponse, error)

	GetUser(string) (*models.UserCRUDResponse, error)

	UpdateUser(string, *models.UserUpdateRequest) (*models.UserCRUDResponse, error)

	UpdateUserPassword(string, *models.UserUpdatePasswordRequest) (*models.UserCRUDResponse, error)

	DeleteUser(string) (*models.UserCRUDResponse, error)
}
