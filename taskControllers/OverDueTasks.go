package controllers

import (
	"net/http"
	"time"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

// GetOverdueTasks godoc
// @Summary      Get user's overdue tasks
// @Description  Returns all overdue (past due date) tasks for the currently logged-in user.
// @Tags         Reports
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}  "List of overdue tasks"
// @Failure      401  {object}  map[string]string        "Unauthorized - invalid or missing token"
// @Failure      500  {object}  map[string]string        "Internal server error"
// @Router       /reports/overdue [get]
func GetOverdueTasks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var tasks []models.Task
	now := time.Now().UTC()

	// Auto-mark overdue tasks
	if err := config.DB.Model(&models.Task{}).
		Where("user_id = ? AND due_date < ? AND status != ?", userID, now, "completed").
		Update("is_overdue", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update overdue tasks"})
		return
	}

	// Now fetch updated overdue tasks
	if err := config.DB.Where("user_id = ? AND is_overdue = ?", userID, true).
		Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch overdue tasks"})
		return
	}

	var response []gin.H

	for _, task := range tasks {
		overdueDays := int(time.Since(task.DueDate).Hours() / 24)

		response = append(response, gin.H{
			"id":           task.Id,
			"title":        task.Title,
			"description":  task.Description,
			"status":       task.Status,
			"is_overdue":   task.IsOverdue,
			"overdue_days": overdueDays,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(response),
		"data":   response,
	})
}
