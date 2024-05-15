package model

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	gorm.Model
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	BirthData time.Time `json:"birth-data"`
	UserID    uint      `json:"user-id"`
}
