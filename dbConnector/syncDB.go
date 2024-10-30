package dbConnector

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func SyncDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = DB.AutoMigrate()
	if err != nil {
		log.Fatal("Error migrating the database")
	}
	// Sync database
	fmt.Println("Database synced")
}
