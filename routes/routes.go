package Routes

import (
	"todo-manager/docs"
	middlewares "todo-manager/middleware"
	taskControllers "todo-manager/taskControllers"
	userControllers "todo-manager/userControllers"

	_ "todo-manager/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// admin routes
	admin := r.Group("/admin")
	admin.Use(middlewares.AuthMiddleware())
	{

		admin.GET("/users", userControllers.AdminGetAllUsers)

		admin.GET("/tasks", userControllers.AdminGetAllTasks)

		admin.GET("/searchTask", userControllers.SearchTasks) //search tasks by title or description

		admin.GET("/tasksByFilter", userControllers.AdminGetTasksByFilter) //get tasks by filter

	}

	// Protected routes
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{

		auth.POST("/task", taskControllers.CreateTask) //create new task

		auth.GET("/tasks", taskControllers.GetAllTasks) //read all task

		auth.GET("/tasks/:id", taskControllers.GetTaskByID) // read one task by id

		auth.GET("/tasks/group/:group", taskControllers.GetTasksByGroup) //read task by group

		auth.GET("/tasksByFilter", taskControllers.GetTasksByFilter) //read task by filter date and status

		auth.POST("/tasks/assign", taskControllers.AssignTask) //assign task by admin or user

		auth.DELETE("/deleteTask/:id", taskControllers.DeleteTask) //delete task by id

		auth.PUT("/updateTasks/:id", taskControllers.UpdateTask) //update task by id

		auth.GET("reports/overdue", taskControllers.GetOverdueTasks)

		auth.GET("reports/tomorrow", taskControllers.GetTomorrowTasks)

		auth.GET("reports/week", taskControllers.GetWeekTasks)

	}

	return r
}
