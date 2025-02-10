package repository

import (
	"errors"
	"gorm.io/gorm"
	ownErrors "movie-reservation-system/errors"
	"movie-reservation-system/models"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func CreateRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) CreateUser(user *models.UserDB) (*models.UserDB, error) {
	result := ur.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepositoryImpl) GetUser(email string) (*models.UserDB, error) {
	user := &models.UserDB{}

	result := ur.db.First(user, "email = ?", email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ownErrors.ErrorUserNotExist{Email: email}
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepositoryImpl) UpdateUser(user *models.UserDB) (*models.UserDB, error) {
	result := ur.db.Save(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepositoryImpl) DeleteUser(user *models.UserDB) (*models.UserDB, error) {
	result := ur.db.Delete(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
