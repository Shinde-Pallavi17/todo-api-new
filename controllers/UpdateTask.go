package controllers

import (
	"net/http"
	"time"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required,max=100"`
	Description string `json:"description" binding:"required"`
	DueDate     string `json:"due_date" binding:"required" example:"yyyy-mm-dd"`               // string from client
	Status      string `json:"status" binding:"omitempty,oneof=pending in_progress completed"` // optional
}

// UpdateTask godoc
// @Summary Update a task
// @Description Replace all fields of a task by its ID.
// Due date must be in YYYY-MM-DD format and will be stored in UTC.
// If no status is provided, it defaults to "pending".
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Param task body UpdateTaskRequest true "Updated Task input"
// @Success 200 {object} models.Task
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 404 {object} map[string]string "Task not found"
// @Failure 500 {object} map[string]string "Failed to update task"
// @Router /updateTasks/{id} [put]
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse due date string → store in UTC
	parsedDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	// Overwrite all fields
	task.Title = req.Title
	task.Description = req.Description
	task.DueDate = parsedDate.UTC() // ensure UTC

	// Set status: if empty, default to "pending"
	if req.Status == "" {
		task.Status = "pending"
	} else {
		task.Status = req.Status
	}

	if err := config.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}
