package controllers

import (
	"net/http"
	"time"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required,max=100"`
	Description string `json:"description" binding:"required"`
	DueDate     string `json:"due_date" binding:"required" example:"yyyy-mm-dd"` // user passes string in YYYY-MM-DD
	Status      string `json:"status" binding:"omitempty,oneof=pending in_progress completed"`
}

// CreateTask godoc
// @Summary Create a new task
// @Description Add a new task (JWT required)
// @Tags tasks
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param task body CreateTaskRequest true "Task input"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /task [post]
func CreateTask(c *gin.Context) {
	var req CreateTaskRequest

	// Validate JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse due date
	parseDueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	//Optional validation: due date must be future
	if parseDueDate.Before(time.Now().UTC()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Due date cannot be in the past"})
		return
	}

	// Get user ID from context (set by middleware)
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	userID := userIDVal.(uint)

	// If no status provided, set default
	status := req.Status
	if status == "" {
		status = "pending"
	}

	// Create task
	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     parseDueDate,
		Status:      status,
		UserID:      userID,
	}

	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	// Return response with all fields (id, created_at, updated_at will be auto-filled by GORM)
	c.JSON(http.StatusCreated, task)
}
