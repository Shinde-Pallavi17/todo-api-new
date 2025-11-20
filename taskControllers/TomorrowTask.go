package controllers

import (
	"net/http"
	"time"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

// GetTomorrowTasks godoc
// @Summary      Get user's tasks due tomorrow
// @Description  Returns all tasks whose due date is exactly tomorrow for the logged-in user.
// @Tags         Reports
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}  "List of tomorrow's tasks"
// @Failure      401  {object}  map[string]string        "Unauthorized - invalid or missing token"
// @Failure      500  {object}  map[string]string        "Internal server error"
// @Router       /reports/tomorrow [get]
func GetTomorrowTasks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var tasks []models.Task
	now := time.Now().UTC()
	tomorrow := now.AddDate(0, 0, 1)

	startOfDay := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24 * time.Hour)

	if err := config.DB.Where(
		"user_id = ? AND due_date >= ? AND due_date < ?",
		userID, startOfDay, endOfDay,
	).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tomorrow's tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(tasks),
		"data":   tasks,
	})
}
