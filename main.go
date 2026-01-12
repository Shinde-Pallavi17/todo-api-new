package main

import (
	"log"
	config "todo-manager/Config"
	routes "todo-manager/Routes"
	"todo-manager/internal/reminders"

	"github.com/joho/godotenv"
)

// @title Todo API
// @version 1.0
// @description Simple Todo API with Gin, GORM, MySQL, Swagger
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	log.Println("Main started")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Initialize DB once
	config.ConnectDB()

	//OPTIONAL: SMTP self-test (env controlled)
	//go utils.RunSMTPStartupTest()

	//start reminder worker
	reminders.StartReminderWorker(config.DB)

	//Setup routes
	r := routes.SetupRouter()

	//Start server
	r.Run(":8080")
}
