package controllers

import (
	"net/http"
	"time"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

// GetWeekTasks godoc
// @Summary      Get user's tasks due in next 7 days
// @Description  Returns all tasks due within the next 7 days for the logged-in user.
// @Tags         Reports
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}  "List of next 7 days tasks"
// @Failure      401  {object}  map[string]string        "Unauthorized - invalid or missing token"
// @Failure      500  {object}  map[string]string        "Internal server error"
// @Router       /reports/week [get]
func GetWeekTasks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var tasks []models.Task
	now := time.Now().UTC()
	nextWeek := now.AddDate(0, 0, 7)

	if err := config.DB.Where(
		"user_id = ? AND due_date >= ? AND due_date <= ?",
		userID, now, nextWeek,
	).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch next 7 days tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(tasks),
		"data":   tasks,
	})
}
