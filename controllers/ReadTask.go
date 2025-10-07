package controllers

import (
	"net/http"
	"strconv"
	"time"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Retrieve all tasks from database (no filter)
// @Tags tasks
// @Produce  json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetAllTasks(c *gin.Context) {
	var tasks []models.Task

	if err := config.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	// Convert all date fields from UTC to IST before sending to user
	loc, _ := time.LoadLocation("Asia/Kolkata")
	for i := range tasks {
<<<<<<< HEAD
=======

		//timestamp converted into utc to ist
>>>>>>> learning
		tasks[i].DueDate = tasks[i].DueDate.In(loc)
		tasks[i].CreatedAt = tasks[i].CreatedAt.In(loc)
		tasks[i].UpdatedAt = tasks[i].UpdatedAt.In(loc)
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID godoc
// @Summary Get a task by ID
// @Description Retrieve a single task by its ID
// @Tags tasks
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [get]
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := config.DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Convert all date fields from UTC to IST before sending to user
	loc, _ := time.LoadLocation("Asia/Kolkata")
	task.DueDate = task.DueDate.In(loc)
	task.CreatedAt = task.CreatedAt.In(loc)
	task.UpdatedAt = task.UpdatedAt.In(loc)

	c.JSON(http.StatusOK, task)
}

// GetTasksByFilter godoc
// @Summary Get all tasks
// @Description Retrieve all tasks with optional filters (status, due_date)
// @Tags tasks
// @Produce  json
// @Param status query string false "Filter by status (pending, in-progress, completed)"
// @Param due_date query string false "Filter by due date (YYYY-MM-DD)"
// @Success 200 {array} models.Task
// @Failure 400 {object} map[string]string "Invalid due_date format"
// @Failure 500 {object} map[string]string "Failed to fetch tasks"
// @Router /tasksByFilter [get]
func GetTasksByFilter(c *gin.Context) {
	var tasks []models.Task

	// Start building query
	query := config.DB.Model(&models.Task{})

	// Filter by status
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Filter by due date
	if dueDateStr := c.Query("due_date"); dueDateStr != "" {
		dueDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due_date format, use YYYY-MM-DD"})
			return
		}
		query = query.Where("DATE(due_date) = ?", dueDate.Format("2006-01-02"))
	}

	// Execute query
	if err := query.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	// Convert all date fields from UTC to IST before sending to user
	loc, _ := time.LoadLocation("Asia/Kolkata")
	for i := range tasks {

		//timestamp converted into utc to ist
		tasks[i].DueDate = tasks[i].DueDate.In(loc)
		tasks[i].CreatedAt = tasks[i].CreatedAt.In(loc)
		tasks[i].UpdatedAt = tasks[i].UpdatedAt.In(loc)
	}

	c.JSON(http.StatusOK, tasks)
}
