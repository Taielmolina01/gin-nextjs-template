package service

import (
	"encoding/base64"
	authErrors "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/auth/errors"
	authUtils "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/auth/utils"
	userErrors "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/errors"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
	userRepository "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/repository"
	userUtils "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/utils"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"math/rand"
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

func NewAuthService(userRepository userRepository.UserRepository, algorithm, secretKey string) AuthService {
	chooseSigningMethod(algorithm, secretKey)

	return &AuthServiceImpl{
		userRepository: userRepository,
	}
}

func (aui *AuthServiceImpl) Login(req *models.UserLoginRequest) (*models.UserLogResponse, error) {
	user, err := aui.userRepository.GetUser(req.Email)
	if err != nil {
		return nil, userErrors.ErrorUserNotExist{Email: req.Email}
	}

	if !userUtils.ValidatePassword(user.Password, req.Password) {
		return nil, userErrors.ErrorWrongOldPassword{}
	}

	expiresAt := time.Now().Add(time.Hour * expiresHours)

	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(signingMethod, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return nil, authErrors.ErrorSigningToken{TypeError: err}
	}

	refreshtoken, err := generateRefreshToken()

	if err != nil {
		return nil, authErrors.ErrorGeneratingRefreshToken{}
	}

	response := authUtils.MapUserDBToLogResponse(user, signedToken, refreshtoken)

	return response, nil
}

func (aui *AuthServiceImpl) Logout(userEmail string) (*models.UserLogResponse, error) {
	user, err := aui.userRepository.GetUser(userEmail)

	if err != nil {
		return nil, userErrors.ErrorUserNotExist{Email: userEmail}
	}

	return authUtils.MapUserDBToLogResponse(user, "", ""), nil
	// deberia pasarle los tokens, medio q si me deslogeo no se pa q los quiero pero bueno
}

func generateRefreshToken() (string, error) {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}
