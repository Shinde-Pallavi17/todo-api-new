package config

import (
	"log"
	"todo-manager/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/task?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run AutoMigrate on all models
	database.AutoMigrate(&models.Task{})

	DB = database
	log.Println("Database connected and migrated successfully")
}
