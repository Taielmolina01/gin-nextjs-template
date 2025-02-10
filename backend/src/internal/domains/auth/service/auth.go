package service

import (
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"math/rand"
	userErrors "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/errors"
	authErrors "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/auth/errors"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
	userRepository "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/repository"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/utils"
	"time"
)

type AuthServiceImpl struct {
	userRepository userRepository.UserRepository
}

var signingMethod jwt.SigningMethod
var secretKey []byte

const expiresHours = 2

func chooseSigningMethod(algorithm string, key string) {
	switch algorithm {
	case "HS256":
		signingMethod = jwt.SigningMethodHS256
	case "RS256":
		signingMethod = jwt.SigningMethodRS256
	default:
		log.Fatalf("Unsupported JWT algorithm: %s", algorithm)
	}
	secretKey = []byte(key)
}

func GetSigningMethod() jwt.SigningMethod {
	return signingMethod
}

func NewAuthService(authRepository authRepository.AuthRepository, userRepository userRepository.UserRepository, algorithm, secretKey string) AuthService {
	chooseSigningMethod(algorithm, secretKey)

	return &AuthServiceImpl{
		userRepository: userRepository,
		authRepository: authRepository,
	}
}


func (aui *AuthServiceImpl) Login(req *models.UserLoginRequest) (*models.UserLoginResponse, error) {
	user, err := aui.userRepository.GetUser(req.Email)
	if err != nil {
		return "", ownErrors.ErrorUserNotExist{Email: req.Email}
	}

	if !utils.ValidatePassword(user.Password, req.Password) {
		return "", ownErrors.ErrorWrongOldPassword{}
	}

	expiresAt := time.Now().Add(time.Hour * expiresHours)

	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(signingMethod, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", ownErrors.ErrorSigningToken{TypeError: err}
	}

	refreshtoken := 

	response := MapUserDBToLoginResponse(user, signedToken, refreshtoken)

	return response, nil
}

func (aui *AuthServiceImpl) Logout(userEmail string) (string, error) {
	_, err := aui.userRepository.GetUser(userEmail)

	if err != nil {
		return "", ownErrors.ErrorUserNotExist{Email: userEmail}
	}

	token, err := aui.authRepository.GetToken(userEmail)

	if err != nil {
		return "", ownErrors.ErrorUserTokenNotExist{UserEmail: userEmail}
	}

	return "user logged succesfully", nil
}

func generateRefreshToken() (string, error) {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}
