package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/repository"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/utils"
	userErrors "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/errors"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (us *UserServiceImpl) CreateUser(req *models.UserRequest) (*models.UserCRUDResponse, error) {

	if req.Role == "" {
		req.Role = "user"
	}

	// Validate fields of request
	if err := utils.ValidateUserFields(req); err != nil {
		return nil, err
	}

	// Call to the db to validate that the user doesnt already exist
	_, userError := us.GetUser(req.Email)

	var userNotExistErr userErrors.ErrorUserNotExist
	if userError != nil && !errors.As(userError, &userNotExistErr) {
		return nil, userError
	}

	// Must hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		return nil, userErrors.ErrorEncriptyngPassword{}
	}

	req.Password = string(hashedPassword)

	newUser := utils.MapUserRequestToUserDB(req)

	// Save user in the db
	return us.UserRepository.CreateUser(newUser)
}

func (us *UserServiceImpl) GetUser(email string) (*models.UserCRUDResponse, error) {
	// Get user from the db
	return us.UserRepository.GetUser(email)
}

func (us *UserServiceImpl) UpdateUser(email string, req *models.UserUpdateRequest) (*models.UserCRUDResponse, error) {
	// Get user from the db
	user, err := us.GetUser(email)

	if err != nil {
		return nil, userErrors.ErrorUserNotExist{Email: email}
	}

	if err := utils.ValidateAndUpdateUser(req, user); err != nil {
		return nil, err
	}

	// Save updated user in the db
	return us.UserRepository.UpdateUser(user)
}

func (us *UserServiceImpl) UpdateUserPassword(email string, req *models.UserUpdatePasswordRequest) (*models.UserCRUDResponse, error) {
	// Get user from the db
	user, err := us.GetUser(email)

	if err != nil {
		return nil, userErrors.ErrorUserNotExist{Email: email}
	}

	// Validate password fields
	if !utils.ValidatePassword(user.Password, req.OldPassword) {
		return nil, userErrors.ErrorWrongOldPassword{}
	}

	if len(req.NewPassword) < 8 {
		return nil, userErrors.ErrorPasswordMustHaveLenght8{}
	}

	// Update password
	user.Password = req.NewPassword // Must hash the password here

	// Save updated user in the db
	return us.UserRepository.UpdateUser(user)
}

func (us *UserServiceImpl) DeleteUser(email string) (*models.UserCRUDResponse, error) {
	// Get user from the db
	user, err := us.GetUser(email)

	if err != nil {
		return nil, userErrors.ErrorUserNotExist{email}
	}

	// Delete user from the db
	return us.UserRepository.DeleteUser(user)
}
