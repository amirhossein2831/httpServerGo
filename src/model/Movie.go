package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name        string `json:"name"`
	Director    string `json:"director"`
	Publication string `json:"publication"`
	WatchTime   int    `json:"watch_time"`
}

func (m Movie) IsModel() {}

func MovieToMod(movies []Movie) []Mod {
	mods := make([]Mod, len(movies))
	for i, u := range movies {
		mods[i] = Mod(u)
	}
	return mods
}
