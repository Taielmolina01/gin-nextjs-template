package repository

import (
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
)

type UserRepository interface {
	CreateUser(*models.UserDB) (*models.UserCRUDResponse, error)

	GetUser(string) (*models.UserCRUDResponse, error)

	UpdateUser(*models.UserDB) (*models.UserCRUDResponse, error)

	DeleteUser(*models.UserDB) (*models.UserCRUDResponse, error)
}
