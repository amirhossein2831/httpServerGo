package model

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
		&Profile{},
		&Book{},
		&Movie{},
		&Role{},
		&Permission{},
		&RolePermission{},
		&UserRole{},
	)

	if err != nil {
		return err
	}
	return nil
}
