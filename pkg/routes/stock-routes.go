package routes

import (
	"github.com/Sambit99/Basic-StocksAPI-Go-Postgres/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterStocksRoutes = func(r *mux.Router) {
	r.HandleFunc("/api/stock", controller.GetAllStocks).Methods("GET")
	r.HandleFunc("/api/stock", controller.CreateStock).Methods("POST")
	r.HandleFunc("/api/stock/{stockId}", controller.GetStockByID).Methods("GET")
	r.HandleFunc("/api/stock/{stockId}", controller.DeleteStock).Methods("DELETE")
	r.HandleFunc("/api/stock/{stockId}", controller.UpdateStock).Methods("PUT")
}
