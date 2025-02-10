package repository

import (
	"errors"
	"gorm.io/gorm"
	ownErrors "movie-reservation-system/errors"
	"movie-reservation-system/models"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (ar *AuthRepositoryImpl) CreateToken(token *models.TokenDB) (*models.TokenDB, error) {
	result := ar.db.Create(&token)

	if result.Error != nil {
		return nil, result.Error
	}

	return token, nil
}

func (ar *AuthRepositoryImpl) GetToken(userEmail string) (*models.TokenDB, error) {
	token := &models.TokenDB{}

	result := ar.db.First(&models.TokenDB{}, "user_email = ?", userEmail)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ownErrors.ErrorUserTokenNotExist{UserEmail: userEmail}
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return token, nil
}

func (ar *AuthRepositoryImpl) DeleteToken(token *models.TokenDB) (*models.TokenDB, error) {
	result := ar.db.Delete(&models.TokenDB{}, "id = ?", token.ID) // Aquí usamos el ID directamente en la condición

	if result.Error != nil {
		return nil, result.Error
	}

	return token, nil
}
