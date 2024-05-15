package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name        string `json:"name"`
	Director    string `json:"director"`
	Publication string `json:"publication"`
	WatchTime   int    `json:"watch_time"`
}
