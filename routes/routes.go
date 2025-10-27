package Routes

import (
	"todo-manager/docs"
	"todo-manager/middlewares"
	taskControllers "todo-manager/taskControllers"
	userControllers "todo-manager/userControllers"

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
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// public routes
	r.POST("/register", userControllers.RegisterUser) //register user

	r.POST("/login", userControllers.LoginUser) //user login

	// Protected routes
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/task", taskControllers.CreateTask) //create new task

		auth.GET("/tasks", taskControllers.GetAllTasks) //read all task

		auth.GET("/tasks/:id", taskControllers.GetTaskByID) // read one task by id

		auth.GET("/tasksByFilter", taskControllers.GetTasksByFilter) //read task by filter date and status

		auth.DELETE("/deleteTask/:id", taskControllers.DeleteTask) //delete task by id

		auth.PUT("/updateTasks/:id", taskControllers.UpdateTask) //update task by id

	}

	return r
}
