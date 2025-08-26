package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Sambit99/Basic-StocksAPI-Go-Postgres/pkg/model"
	"github.com/Sambit99/Basic-StocksAPI-Go-Postgres/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllStocks(w http.ResponseWriter, r *http.Request) {

	stocks := model.GetAllStocks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stocks)
}

func GetStockByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stockId := params["stockId"]
	ID, err := strconv.ParseInt(stockId, 0, 0)
	if err != nil {
		log.Fatal("Error while parsing stock id")
	}

	stock := model.GetStockByID(ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stock)

}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	stock := &model.Stock{}

	utils.ParseBody(r, stock)

	s := stock.CreateStock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stockId := params["stockId"]

	fmt.Println(stockId)

	ID, err := strconv.ParseInt(stockId, 0, 0)
	if err != nil {
		log.Fatal("Error while parsing stock id")
	}

	stock := model.DeleteStock(ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stock)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stockId := params["stockId"]

	ID, err := strconv.ParseInt(stockId, 0, 0)

	if err != nil {
		log.Fatal("Error while parsing stock id")
	}

	var stock model.Stock

	utils.ParseBody(r, &stock)
	stock.ID = uint(ID)

	s := model.UpdateStock(stock)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}
