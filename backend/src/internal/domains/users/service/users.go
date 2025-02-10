package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	ownErrors "movie-reservation-system/errors"
	"movie-reservation-system/models"
	"movie-reservation-system/repository/user"
	"movie-reservation-system/service"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func mapUserRequestToUserDB(req *models.UserRequest) *models.UserDB {
	return &models.UserDB{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Role:     req.Role,
	}
}

func (us *UserServiceImpl) CreateUser(req *models.UserRequest) (*models.UserDB, error) {

	if req.Role == "" {
		req.Role = "user"
	}

	// Validate fields of request
	if err := service.ValidateUserFields(req); err != nil {
		return nil, err
	}

	// Call to the db to validate that the user doesnt already exist
	_, userError := us.GetUser(req.Email)

	var userNotExistErr ownErrors.ErrorUserNotExist
	if userError != nil && !errors.As(userError, &userNotExistErr) {
		return nil, userError
	}

	// Must hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		return nil, ownErrors.ErrorEncriptyngPassword{}
	}

	req.Password = string(hashedPassword)

	newUser := mapUserRequestToUserDB(req)

	// Save user in the db
	return us.UserRepository.CreateUser(newUser)
}

func (us *UserServiceImpl) GetUser(email string) (*models.UserDB, error) {
	// Get user from the db
	return us.UserRepository.GetUser(email)
}

func (us *UserServiceImpl) UpdateUser(email string, req *models.UserUpdateRequest) (*models.UserDB, error) {
	// Get user from the db
	user, err := us.GetUser(email)

	if err != nil {
		return nil, ownErrors.ErrorUserNotExist{Email: email}
	}

	if err := service.ValidateAndUpdateUser(req, user); err != nil {
		return nil, err
	}

	// Save updated user in the db
	return us.UserRepository.UpdateUser(user)
}

func (us *UserServiceImpl) UpdateUserPassword(email string, req *models.UserUpdatePasswordRequest) (*models.UserDB, error) {
	// Get user from the db
	user, err := us.GetUser(email)

	if err != nil {
		return nil, ownErrors.ErrorUserNotExist{Email: email}
	}

	// Validate password fields
	if !service.ValidatePassword(user.Password, req.OldPassword) {
		return nil, ownErrors.ErrorWrongOldPassword{}
	}

	if len(req.NewPassword) < 8 {
		return nil, ownErrors.ErrorPasswordMustHaveLenght8{}
	}

	// Update password
	user.Password = req.NewPassword // Must hash the password here

	// Save updated user in the db
	return us.UserRepository.UpdateUser(user)
}

func (us *UserServiceImpl) DeleteUser(email string) (*models.UserDB, error) {
	// Get user from the db
	user, err := us.GetUser(email)

	if err != nil {
		return nil, ownErrors.ErrorUserNotExist{email}
	}

	// Delete user from the db
	return us.UserRepository.DeleteUser(user)
}
