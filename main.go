package main

import (
	config "todo-manager/Config"
	routes "todo-manager/Routes"
)

// @title Todo API
// @version 1.0
// @description Simple Todo API with Gin, GORM, MySQL, Swagger
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize DB once
	config.ConnectDB()

	// Setup routes
	r := routes.SetupRouter()

	// Start server
	r.Run(":8080")
}
