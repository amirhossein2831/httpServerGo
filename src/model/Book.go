package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(25);not null" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b Book) IsModel() {}

func BookToMod(books []Book) []Mod {
	mods := make([]Mod, len(books))
	for i, u := range books {
		mods[i] = Mod(u)
	}
	return mods
}
