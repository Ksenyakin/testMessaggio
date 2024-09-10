package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
)

var producer sarama.SyncProducer

func InitProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	var err error
	producer, err = sarama.NewSyncProducer([]string{os.Getenv("KAFKA_BROKER")}, config)
	if err != nil {
		log.Fatal("Failed to start Kafka producer:", err)
	}
}

func CloseProducer() {
	if err := producer.Close(); err != nil {
		log.Println("Failed to close Kafka producer:", err)
	}
}

func SendMessage(message string) error {
	msg := &sarama.ProducerMessage{
		Topic: os.Getenv("KAFKA_TOPIC"),
		Value: sarama.StringEncoder(message),
	}

	_, _, err := producer.SendMessage(msg)
	return err
}
