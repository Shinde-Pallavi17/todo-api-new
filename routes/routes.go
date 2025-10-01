package routes

import (
	"todo-manager/controllers"
	"todo-manager/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "todo-manager/docs"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Swagger setup
	docs.SwaggerInfo.Title = "Todo API"
	docs.SwaggerInfo.Description = "Simple Todo API with Gin, GORM, MySQL, Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	r.POST("/task", controllers.CreateTask) //create new task

	r.GET("/tasks", controllers.GetAllTasks) //read all task

	r.GET("/tasks/:id", controllers.GetTaskByID) // read one

	r.GET("/tasksByFilter", controllers.GetTasksByFilter) //read all task

	r.DELETE("/deleteTask/:id", controllers.DeleteTask) //delete task by id

	r.PUT("/updateTasks/:id", controllers.UpdateTask) //update task by id

	return r
}
