package model

import (
	config "github.com/Sambit99/Basic-StocksAPI-Go-Postgres/pkg/Config"
	"gorm.io/gorm"
)

type Stock struct {
	*gorm.Model
	Name    string  `json:"name"`
	Price   float64 `gorm:"type:numeric(10,2)" json:"price"`
	Company string  `json:"company"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Stock{})
}

func GetAllStocks() []Stock {
	var stocks []Stock
	db.Find(&stocks)
	return stocks
}
