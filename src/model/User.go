package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string  `json:"first-name"`
	LastName  string  `json:"last-name"`
	Email     string  `gorm:"unique;not null" json:"email"`
	Password  string  `json:"password"`
	Profile   Profile `json:"profile"`
	Roles     []Role  `gorm:"many2many:user_roles;" json:"roles"`
}

type UserRole struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`
}
