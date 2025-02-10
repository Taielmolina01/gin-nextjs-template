package utils

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/errors"
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
	"reflect"
	"strings"
)

func trimStructFields(s interface{}) {
	val := reflect.ValueOf(s)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldName := val.Type().Field(i).Name

			if field.Kind() == reflect.String && fieldName != "Password" {
				field.SetString(strings.TrimSpace(field.String()))
			}
		}
	}
}

func ValidateUserFields(req *models.UserRequest) error {
	trimStructFields(req)

	if req.Email == "" {
		return errors.ErrorUserMustHaveEmail{}
	}

	if req.Name == "" {
		return errors.ErrorUserMustHaveName{}
	}

	if len(req.Password) < 8 {
		return errors.ErrorPasswordMustHaveLenght8{}
	}

	if !Contains(models.GetRoles(), string(req.Role)) {
		return errors.ErrorUserRoleInvalid{Role: string(req.Role)}
	}

	return nil
}

func ValidateAndUpdateUser(req *models.UserUpdateRequest, user *models.UserDB) error {
	trimStructFields(req)

	// Validate fields
	if req.Name != nil && *req.Name == "" {
		return errors.ErrorUserMustHaveName{}
	}

	if req.Role != nil && !Contains(models.GetRoles(), string(*req.Role)) {
		return errors.ErrorUserRoleInvalid{Role: string(*req.Role)}
	}

	// Update fileds
	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.Role != nil {
		user.Role = *req.Role
	}

	return nil
}

func Contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func ValidatePassword(storedPassword string, enteredPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(enteredPassword)) == nil
}

func MapUserRequestToUserDB(req *models.UserRequest) *models.UserDB {
	return &models.UserDB{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Role:     req.Role,
	}
}