package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"myservice/internal/models"
	"os"
)

var db *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	db.Close()
}

func SaveMessage(message models.Message) error {
	_, err := db.Exec("INSERT INTO messages (content, processed) VALUES ($1, $2)", message.Content, false)
	return err
}

func MarkMessageAsProcessed(id int) error {
	_, err := db.Exec("UPDATE messages SET processed = true WHERE id = $1", id)
	return err
}

func GetStatistics() (map[string]int, error) {
	var total, processed int
	err := db.QueryRow("SELECT COUNT(*) FROM messages").Scan(&total)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow("SELECT COUNT(*) FROM messages WHERE processed = true").Scan(&processed)
	if err != nil {
		return nil, err
	}

	return map[string]int{"total": total, "processed": processed}, nil
}
