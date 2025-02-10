package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TokenDB struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	AccessToken  string    `gorm:"type:varchar(255);not null"`
	RefreshToken string    `gorm:"type:varchar(255);not null"`
	UserEmail    string    `gorm:"type:varchar(255);not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	User         UserDB    `gorm:"foreignKey:UserEmail;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	gorm.Model
}

type TokenResponse struct {
    AccessToken struct {
        ID           string    `json:"ID"`
        AccessToken  string    `json:"AccessToken"`
        RefreshToken string    `json:"RefreshToken"`
        UserEmail    string    `json:"UserEmail"`
        ExpiresAt    time.Time `json:"ExpiresAt"`
        User         struct {
            Email    string `json:"Email"`
            Name     string `json:"Name"`
            Password string `json:"Password"`
            Role     string `json:"Role"`
            ID       int    `json:"ID"`
        } `json:"User"`
        CreatedAt time.Time `json:"CreatedAt"`
        UpdatedAt time.Time `json:"UpdatedAt"`
    } `json:"access_token"`
}