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
	Description string `json:"description"`
	DueDate     string `json:"due_date" example:"yyyy-mm-dd" gorm:"optional"` // user passes string in YYYY-MM-DD
	ReminderAt  string `json:"reminder_at" example:"2025-01-02T10:00:00+05:30"`
	Priority    string `json:"priority" binding:"omitempty,oneof=medium high low"`
	TaskGroup   string `json:"task_group" binding:"oneof=personal office shopping family friends education health travel food"`
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

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parseDueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	//Normalize today's date to midnight UTC
	today := time.Now().UTC().Truncate(24 * time.Hour)

	isPast := parseDueDate.Before(today)
	if isPast {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Due date cannot be in the past"})
		return
	}

	var reminderTime *time.Time

	if req.ReminderAt != "" {
		parsedReminder, err := time.Parse(time.RFC3339, req.ReminderAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reminder format (use RFC3339)"})
			return
		}

		if parsedReminder.Before(time.Now().UTC()) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Reminder time must be in the future"})
			return
		}

		reminderTime = &parsedReminder
	}

	//Get user ID from context (set by middleware)
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	userID := userIDVal.(uint)

	status := req.Status
	if status == "" {
		status = "pending"
	}

	priority := req.Priority
	if priority == "" {
		priority = "medium"
	}

	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     parseDueDate,
		ReminderAt:  reminderTime,
		Priority:    priority,
		TaskGroup:   req.TaskGroup,
		Status:      status,
		UserID:      userID,
	}

	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	//Return response with all fields (id, created_at, updated_at will be auto-filled by GORM)
	c.JSON(http.StatusCreated, task)

	if req.ReminderAt != "" {
		reminderTime, err := time.Parse(time.RFC3339, req.ReminderAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reminder time format"})
			return
		}

		reminder := models.Reminder{
			TaskID:     task.Id,
			UserID:     userID,
			ReminderAt: reminderTime.UTC(),
		}

		config.DB.Create(&reminder)
	}

}
