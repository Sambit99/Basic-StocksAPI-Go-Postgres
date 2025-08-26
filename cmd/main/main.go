package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Sambit99/Basic-StocksAPI-Go-Postgres/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterStocksRoutes(r)

	var srv = &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server running on port: 8080")

	log.Fatal(srv.ListenAndServe())
}
