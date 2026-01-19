package controllers

import (
	"time"
	config "todo-manager/Config"
	"todo-manager/models"
	"todo-manager/utils"

	"github.com/gin-gonic/gin"
)

type AssignTaskRequest struct {
	Title         string  `json:"title" binding:"required"`
	Description   *string `json:"description"`
	AssignToEmail string  `json:"assign_to_email" binding:"required,email"`

	DueDate    string `json:"due_date" example:"yyyy-mm-dd" gorm:"optional"`
	ReminderAt string `json:"reminder_at" example:"2025-01-02T10:00:00+05:30"`

	Priority  string `json:"priority"`
	Status    string `json:"status"`
	TaskGroup string `json:"task_group"`
}

// AssignTask godoc
// @Summary      Assign a task to a user
// @Description  Admin can assign a task to any user. A normal user can assign a task only to another user (not admin).
// @Tags         tasks
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        task body AssignTaskRequest true "Task assignment payload"
// @Success      201 {object} map[string]interface{} "Task assigned successfully"
// @Failure      400 {object} map[string]string "Invalid request body"
// @Failure      401 {object} map[string]string "Unauthorized"
// @Failure      403 {object} map[string]string "Permission denied"
// @Failure      404 {object} map[string]string "Assignee not found"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /tasks/assign [post]
func AssignTask(c *gin.Context) {

	//Read request
	var req AssignTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//Get assigner from JWT
	assignerID := c.GetUint("userID")

	assignerName := c.GetString("username")

	if assignerName == "" {
		var assigner models.User
		if err := config.DB.First(&assigner, assignerID).Error; err == nil {
			assignerName = assigner.Username
		}
	}

	assignerRole := c.GetString("role")

	var assignee models.User
	if err := config.DB.Where("email = ?", req.AssignToEmail).First(&assignee).Error; err != nil {
		c.JSON(404, gin.H{"error": "Assignee not found"})
		return
	}

	if assignerRole == "user" && assignee.Role == "admin" {
		c.JSON(403, gin.H{"error": "User cannot assign task to admin"})
		return
	}

	// safe description handling
	desc := ""
	if req.Description != nil {
		desc = *req.Description
	}

	dueDate := parseDate(req.DueDate)
	reminderAt := parseDateTime(req.ReminderAt)

	// inline defaults instead of undefined defaultVal
	priority := req.Priority
	if priority == "" {
		priority = "medium"
	}
	status := req.Status
	if status == "" {
		status = "pending"
	}
	taskGroup := req.TaskGroup
	if taskGroup == "" {
		taskGroup = "personal"
	}

	var reminderAtPtr *time.Time
	if !reminderAt.IsZero() {
		reminderAtPtr = &reminderAt
	}

	task := models.Task{
		Title:       req.Title,
		Description: desc,
		UserID:      assignee.ID,

		DueDate:    dueDate,
		ReminderAt: reminderAtPtr,
		Priority:   priority,
		Status:     status,
		TaskGroup:  taskGroup,

		AssignedByID: assignerID,
		AssignedBy:   assignerName,
		AssignedRole: assignerRole,
	}

	// Save task
	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to assign task"})
		return
	}

	//SEND EMAIL IMMEDIATELY AFTER ASSIGNMENT
	_ = utils.SendTaskAssignedEmail(
		assignee.Email,
		assignee.Username,
		task,
		assignerName,
		assignerRole,
	)
	c.JSON(201, gin.H{
		"message": "Task assigned successfully",
		"task":    task,
	})
}

// parseDate parses a date in "2006-01-02" or RFC3339 and returns zero time on error/empty.
func parseDate(s string) time.Time {
	if s == "" {
		return time.Time{}
	}
	if t, err := time.Parse("2006-01-02", s); err == nil {
		return t
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t
	}
	return time.Time{}
}

// parseDateTime parses a datetime in RFC3339 or "2006-01-02 15:04:05" and returns zero time on error/empty.
func parseDateTime(s string) time.Time {
	if s == "" {
		return time.Time{}
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t
	}
	if t, err := time.Parse("2006-01-02 15:04:05", s); err == nil {
		return t
	}
	return time.Time{}
}
