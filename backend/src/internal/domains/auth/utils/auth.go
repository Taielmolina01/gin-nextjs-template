package utils

import (
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
)

func MapUserDBToLoginResponse(user *models.UserDB, token string) *models.UserLoginResponse {
	return &models.UserLoginResponse{
		Email: user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Role: user.Role,
		Token: token,
	}
}