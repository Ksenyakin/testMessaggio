package services

import (
	"myservice/internal/db"
	"myservice/internal/kafka"
	"myservice/internal/models"
)

func SaveAndSendMessage(message models.Message) error {
	if err := db.SaveMessage(message); err != nil {
		return err
	}

	if err := kafka.SendMessage(message.Content); err != nil {
		return err
	}

	return nil
}
