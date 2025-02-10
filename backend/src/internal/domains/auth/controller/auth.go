package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	ownErrors "movie-reservation-system/errors"
	"movie-reservation-system/models"
	"movie-reservation-system/service/auth"
	"net/http"
)

type AuthController struct {
	AuthService auth.AuthService
}

func NewAuthController(authService auth.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var request models.UserLoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	token, err := ac.AuthService.Login(&request)

	if err != nil {
		if errors.Is(err, ownErrors.ErrorUserNotExist{Email: request.Email}) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}

func (ac *AuthController) Logout(ctx *gin.Context) {
	email := ctx.Param("email")

	token, err := ac.AuthService.Logout(email)

	if err != nil {
		if errors.Is(err, ownErrors.ErrorUserNotExist{Email: email}) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else if errors.Is(err, ownErrors.ErrorUserTokenNotExist{UserEmail: email}) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": token,
	})
}
