package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	Host     string          `json:"host"`
	URL      string          `json:"url"`
	Method   string          `json:"method"`
	Protocol string          `json:"protocol"`
	Header   json.RawMessage `json:"header"`
	Body     json.RawMessage `json:"body"`
}
