package api

import (
	"encoding/json"
	"net/http"

	"myservice/internal/db"
	"myservice/internal/models"
	"myservice/internal/services"
)

func CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.SaveAndSendMessage(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetStatisticsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := db.GetStatistics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stats)
}
