package models

import (
	"gorm.io/gorm"
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
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Role     Role   `json:"role" binding:"omitempty,oneof=user admin"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
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
	Email    string `gorm:"type:varchar(100);primaryKey"`
	Name     string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null" validate:"required,min=8"`
	Role     Role   `gorm:"type:varchar(30);default:user`
	gorm.Model
}
