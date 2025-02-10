package middlewares

import (
	"fmt"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/auth/service"
	userErrors "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/errors"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	JWTSecretKey string
}

func NewAuthMiddleware(secretKey string) *AuthMiddleware {
	return &AuthMiddleware{JWTSecretKey: secretKey}
}

func (a *AuthMiddleware) commonAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")

		var user models.UserDB
		if err := db.First(&user, "email = ?", email).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": userErrors.ErrorUserNotExist{Email: email}.Error()})
			c.Abort()
			return
		}

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
			c.Abort()
			return
		}

		// Verifiy the jwt format
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			c.Abort()
			return
		}

		tokenString = tokenParts[1]

		// Verify if the token is valid
		claims, err := verifyToken(tokenString, a.JWTSecretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			c.Abort()
			return
		}

		userEmail, ok := claims["email"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		if userEmail != email {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Emails don't match"})
			c.Abort()
			return
		}

		c.Set("userEmail", user.Email)
		c.Set("userRole", user.Role)

		c.Next()
	}
}

func (a *AuthMiddleware) PublicAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return a.commonAuthMiddleware(db)
}

func (a *AuthMiddleware) AdminAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		a.commonAuthMiddleware(db)(c)

		if c.IsAborted() {
			return
		}

		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			c.Abort()
			return
		}

		userDB, ok := user.(models.UserDB)
		if !ok || userDB.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func verifyToken(tokenString, secretKey string) (jwt.MapClaims, error) {
	signingMethod := service.GetSigningMethod()
	key := []byte(secretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != signingMethod {
			return nil, fmt.Errorf("Invalid signing method: %v", token.Method.Alg())
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
