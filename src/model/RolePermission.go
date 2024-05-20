package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);not null" json:"name"`
	Description string `json:"description"`

	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type Permission struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);not null" json:"name"`
	Description string `json:"description"`

	Roles []Role `gorm:"many2many:role_permissions;" json:"roles"`
}

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}
