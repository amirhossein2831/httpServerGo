package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(50);not null" json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type Permission struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);not null" json:"name"`
	Description string `json:"description"`

	Roles []Role `gorm:"many2many:role_permissions;" json:"roles"`
}

func (r Role) IsModel()       {}
func (p Permission) IsModel() {}

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}

func HasPermission(user User, permissions []string) bool {
	for _, role := range user.Roles {
		for _, rolePermission := range role.Permissions {
			for _, permission := range permissions {
				if rolePermission.Name == permission {
					return true
				}
			}
		}
	}

	return false
}

func HasRole(user User, roles []string) bool {
	for _, role := range user.Roles {
		for _, strRole := range roles {
			if role.Name == strRole {
				return true
			}
		}
	}
	return false
}

func RoleToMod(roles []Role) []Mod {
	mods := make([]Mod, len(roles))
	for i, u := range roles {
		mods[i] = Mod(u)
	}
	return mods
}

func PermissionToMod(permissions []Permission) []Mod {
	mods := make([]Mod, len(permissions))
	for i, u := range permissions {
		mods[i] = Mod(u)
	}
	return mods
}
