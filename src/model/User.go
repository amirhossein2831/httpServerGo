package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string  `json:"first-name"`
	LastName  string  `json:"last-name"`
	Email     string  `gorm:"unique;not null" json:"email"`
	Profile   Profile `json:"profile"`
}
