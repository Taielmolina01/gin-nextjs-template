package service

import (
	"errors"
	userErrors "github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/errors"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/repository"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(UserRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: UserRepository}
}

func mapUserDBToUserCRUDResponse(user *models.UserDB) *models.UserCRUDResponse {
	return &models.UserCRUDResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}
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
	savedUser, err2 := us.UserRepository.CreateUser(newUser)
	if err2 != nil {
		return nil, err2
	}

	return mapUserDBToUserCRUDResponse(savedUser), nil
}

func (us *UserServiceImpl) GetUser(email string) (*models.UserCRUDResponse, error) {
	// Get user from the db
	foundedUser, err := us.UserRepository.GetUser(email)
	if err != nil {
		return nil, err
	}
	return mapUserDBToUserCRUDResponse(foundedUser), nil
}

func (us *UserServiceImpl) UpdateUser(email string, req *models.UserUpdateRequest) (*models.UserCRUDResponse, error) {
	// Get user from the db
	user, err := us.UserRepository.GetUser(email)

	if err != nil {
		return nil, userErrors.ErrorUserNotExist{Email: email}
	}

	if err := utils.ValidateAndUpdateUser(req, user); err != nil {
		return nil, err
	}

	// Save updated user in the db
	updatedUser, err2 := us.UserRepository.UpdateUser(user)

	if err2 != nil {
		return nil, err2
	}

	return mapUserDBToUserCRUDResponse(updatedUser), nil
}

func (us *UserServiceImpl) UpdateUserPassword(email string, req *models.UserUpdatePasswordRequest) (*models.UserCRUDResponse, error) {
	// Get user from the db
	user, err := us.UserRepository.GetUser(email)

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
	updatedUser, err2 := us.UserRepository.UpdateUser(user)

	if err2 != nil {
		return nil, err2
	}

	return mapUserDBToUserCRUDResponse(updatedUser), nil
}

func (us *UserServiceImpl) DeleteUser(email string) (*models.UserCRUDResponse, error) {
	// Get user from the db
	user, err := us.UserRepository.GetUser(email)

	if err != nil {
		return nil, userErrors.ErrorUserNotExist{email}
	}

	// Delete user from the db
	user, err2 := us.UserRepository.DeleteUser(user)

	if err2 != nil {
		return nil, err2
	}

	return mapUserDBToUserCRUDResponse(user), nil
}
