package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Role string

const (
	User  Role = "user"
	Admin Role = "admin"
)

func GetRoles() []string {
	return []string{"user", "admin"}
}

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	FirstName string `json:"firstname" binding:"required"`	
	LastName  string `json:"lastname" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Role     Role   `json:"role" binding:"omitempty,oneof=user admin"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserLoginResponse struct {
	Email    string `json:"email"`
	FirstName string `json:"firstname"`	
	LastName  string `json:"lastname"`
	Role     Role   `json:"role"`
	Token string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}

type UserUpdateRequest struct {
	Name *string `json:"name" binding:"omitempty"`
	Role *Role   `json:"role" binding:"omitempty,oneof=user admin"`
}

type UserUpdatePasswordRequest struct {
	OldPassword string `json:"oldpassword" binding:"required,min=8"`
	NewPassword string `json:"newpassword" binding:"required,min=8"`
}

type UserDB struct {
	ID	uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email    string `gorm:"type:varchar(100);unique"`
	FirstName string `gorm:"type:varchar(50);not null`
	LastName     string `gorm:"type:varchar(50);not null"`
	Password string `gorm:"type:varchar(100);not null" validate:"required,min=8"`
	Role     Role   `gorm:"type:varchar(30);default:user`
	gorm.Model
}

type UserCRUDResponse struct {
	Email	string	`json:"email"`
	FirstName	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
	Role	Role	`json:"role"`
}
