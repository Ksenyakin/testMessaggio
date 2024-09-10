package main

import (
	"github.com/gorilla/mux"
	"log"
	"myservice/internal/api"
	"myservice/internal/db"
	"myservice/internal/kafka"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Инициализация базы данных
	db.InitDB()
	defer db.CloseDB()

	// Инициализация Kafka
	kafka.InitProducer()
	defer kafka.CloseProducer()

	// Роуты
	router.HandleFunc("/messages", api.CreateMessageHandler).Methods("POST")
	router.HandleFunc("/statistics", api.GetStatisticsHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
