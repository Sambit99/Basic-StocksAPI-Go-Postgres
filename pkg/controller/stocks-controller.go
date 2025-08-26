package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Sambit99/Basic-StocksAPI-Go-Postgres/pkg/model"
)

func GetAllStocks(w http.ResponseWriter, r *http.Request) {

	stocks := model.GetAllStocks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stocks)
}
