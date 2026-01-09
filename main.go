package main

import (
	"log"
	config "todo-manager/Config"
	routes "todo-manager/Routes"
	"todo-manager/internal/reminders"
	"todo-manager/utils"
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
	//Initialize DB once
	config.ConnectDB()

	go func() {
		err := utils.SendReminderEmail(
			"pallavishinde622@gmail.com",
			"SMTP Test",
			"If you receive this, SMTP works",
		)
		if err != nil {
			log.Println("SMTP test failed:", err)
		} else {
			log.Println("SMTP test success")
		}
	}()

	//start reminder worker
	reminders.StartReminderWorker(config.DB)

	//Setup routes
	r := routes.SetupRouter()

	//Start server
	r.Run(":8080")
}
