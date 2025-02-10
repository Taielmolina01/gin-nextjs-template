package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"movie-reservation-system/models"
    "strings"
    "gorm.io/gorm"
    "fmt"
	"github.com/golang-jwt/jwt/v5"
    "movie-reservation-system/service/auth"
    "movie-reservation-system/configuration"
    "movie-reservation-system/errors"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        email := c.Param("email")

        var user models.UserDB
        if err := db.First(&user, "email = ?", email).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrorUserNotExist{Email: email}.Error()})
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
        claims, err := verifyToken(tokenString)
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
            c.JSON(http.StatusBadRequest, gin.H{"error": "Emails dont match"})
            c.Abort()
            return
        }

        c.Set("userEmail", user.Email)
        c.Set("userRole", user.Role)

        c.Next()
    }
}

func verifyToken(tokenString string) (jwt.MapClaims, error) {
    signingMethod := auth.GetSigningMethod()
    secretKey := []byte(configuration.GetConfiguration().JwtSecretKey)

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if token.Method != signingMethod {
            return nil, fmt.Errorf("Invalid signing method: %v", token.Method.Alg())
        }
        return secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("Invalid token")
}