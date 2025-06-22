package models

import "gorm.io/gorm"

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	PasswordHash string         `json:"-"`
	Role         string         `json:"role"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
