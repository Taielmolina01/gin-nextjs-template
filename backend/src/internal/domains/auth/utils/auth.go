package utils

import (
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
)

func MapUserDBToLogResponse(user *models.UserDB, token string) *models.UserLogResponse {
	return &models.UserLogResponse{
		Email: user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Role: user.Role,
		Token: token,
	}
}