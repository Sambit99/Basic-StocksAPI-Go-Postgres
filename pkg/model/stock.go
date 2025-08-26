package model

import (
	config "github.com/Sambit99/Basic-StocksAPI-Go-Postgres/pkg/Config"
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
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

func GetStockByID(stockId int64) Stock {
	var stock Stock
	db.Where("ID=?", stockId).Find(&stock)
	return stock
}

func (s *Stock) CreateStock() *Stock {
	db.Create(s)
	return s
}

func DeleteStock(stockId int64) Stock {
	var stock Stock
	db.Where("ID=?", stockId).Delete(&stock)
	return stock
}

func UpdateStock(updatedStock Stock) Stock {

	var existingStock Stock
	db.Where("ID=?", updatedStock.ID).Find(&existingStock)

	if updatedStock.Name != "" {
		existingStock.Name = updatedStock.Name
	}
	if updatedStock.Price != 0 {
		existingStock.Price = updatedStock.Price
	}
	if updatedStock.Company != "" {
		existingStock.Company = updatedStock.Company
	}

	db.Save(&existingStock)

	return existingStock
}
