package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	userErrors "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/errors"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/auth/service"
	"net/http"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
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

	response, err := ac.AuthService.Login(&request)

	if err != nil {
		if errors.Is(err, userErrors.ErrorUserNotExist{Email: request.Email}) {
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

	session := sessions.Default(ctx)

	session.Set(user.Email, response) 
	if err2 := session.Save(); err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
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
		if errors.Is(err, userErrors.ErrorUserNotExist{Email: email}) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else if errors.Is(err, userErrors.ErrorUserTokenNotExist{UserEmail: email}) {
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

	session := sessions.Default(ctx)

	session.Delete(email)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": token,
	})
}
