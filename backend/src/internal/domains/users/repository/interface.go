package repository

import (
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
)

type UserRepository interface {
	CreateUser(*models.UserDB) (*models.UserDB, error)

	GetUser(string) (*models.UserDB, error)

	UpdateUser(*models.UserDB) (*models.UserDB, error)

	DeleteUser(*models.UserDB) (*models.UserDB, error)
}
