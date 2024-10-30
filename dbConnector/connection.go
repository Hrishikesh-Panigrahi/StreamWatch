package dbConnector

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	var err error
	dbUser := os.Getenv("dbUser")
	dbPassword := os.Getenv("dbPassword")
	dbName := os.Getenv("dbName")
	dbTcp := os.Getenv("dbTcp")
	dsn := dbUser + ":" + dbPassword + "@" + dbTcp + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected")
}
