package controllers

import (
	"net/http"
	"time"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

// AdminGetTasksByFilter godoc
// @Summary Get all tasks  (Admin only)
// @Description Retrieve all tasks with optional filters (priority, status, due_date) (JWT required)
// @Tags Admin
// @Security BearerAuth
// @Produce  json
// @Param priority query string false "Filter by priority (high, medium, low)"
// @Param status query string false "Filter by status (pending, in-progress, completed)"
// @Param due_date query string false "Filter by due date (YYYY-MM-DD)"
// @Success 200 {array} models.Task
// @Failure 400 {object} map[string]string "Invalid due_date format"
// @Failure 500 {object} map[string]string "Failed to fetch tasks"
// @Router /admin/tasksByFilter [get]
func AdminGetTasksByFilter(c *gin.Context) {
	// Admin-only check
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "admin access only"})
		return
	}

	var tasks []models.Task

	userID := c.GetUint("userID") // Only current user's tasks

	//Start building query
	query := config.DB.Model(&models.Task{}).Where("user_id = ?", userID)

	//Filter by priority
	if priority := c.Query("priority"); priority != "" {
		query = query.Where("priority = ?", priority)
	}

	//Filter by status
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	//Filter by due date
	if dueDateStr := c.Query("due_date"); dueDateStr != "" {
		dueDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due_date format, use YYYY-MM-DD"})
			return
		}
		query = query.Where("DATE(due_date) = ?", dueDate.Format("2006-01-02"))
	}

	//Execute query
	if err := query.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	//Convert all date fields from UTC to IST before sending to user
	loc, _ := time.LoadLocation("Asia/Kolkata")
	for i := range tasks {

		//timestamp converted into utc to ist
		tasks[i].DueDate = tasks[i].DueDate.In(loc)
		tasks[i].CreatedAt = tasks[i].CreatedAt.In(loc)
		tasks[i].UpdatedAt = tasks[i].UpdatedAt.In(loc)
	}

	c.JSON(http.StatusOK, tasks)
}
