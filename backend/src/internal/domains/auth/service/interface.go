package service

import (
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
)

type AuthService interface {
	Login(*models.UserLoginRequest) (*models.UserLoginResponse, error)

	Logout(string) (*models.TokenDB, error)
}
