package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string  `json:"first-name"`
	LastName  string  `json:"last-name"`
	Email     string  `gorm:"uniqueIndex;not null" json:"email"`
	Password  string  `json:"password"`
	Profile   Profile `json:"profile"`
}
