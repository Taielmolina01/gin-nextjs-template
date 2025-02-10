package service

import (
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
)

type AuthService interface {
	Login(*models.UserLoginRequest) (*models.UserLogResponse, error)

	Logout(string) (*models.UserLogResponse, error)
}
