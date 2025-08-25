package model

import (
	"gorm.io/gorm"
)

type Stock struct {
	*gorm.Model
	Name    string  `json:"name"`
	Price   float64 `gorm:"type:numeric(10,2)" json:"price"`
	Company string  `json:"company"`
}
