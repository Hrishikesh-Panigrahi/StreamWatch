package dbConnector

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	models "github.com/Hrishikesh-Panigrahi/StreamWatch/models"
)

func SyncDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = DB.AutoMigrate(&models.User{}, &models.Video{}, &models.Likes{}, &models.WatchLog{}, &models.Dislikes{})
	if err != nil {
		log.Fatal("Error migrating the database")
	}
	fmt.Println("Database Synced")
}
