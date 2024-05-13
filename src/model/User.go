package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Profile struct {
	gorm.Model
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	BirthData time.Time `json:"birth-data"`
	UserID    uint      `gorm:"foreignKey:UserID" json:"user-id"` // Foreign key

}

type User struct {
	gorm.Model
	FirstName string  `json:"first-name"`
	LastName  string  `json:"last-name"`
	Email     string  `gorm:"unique;not null" json:"email"`
	Profile   Profile `gorm:"foreignKey:UserID" json:"profile"`
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("user_id = ?", u.ID).Delete(&Profile{})
	return
}
