package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Response struct {
	gorm.Model
	Status    int             `json:"status"`
	IsSuccess bool            `json:"isSuccess"`
	Body      json.RawMessage `json:"body"`
}
