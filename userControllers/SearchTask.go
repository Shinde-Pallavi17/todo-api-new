package controllers

import (
	"net/http"
	"strings"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

// SearchTasks godoc
// @Summary Search tasks (Admin only)
// @Description Admin-only API to search tasks by title or description
// @Tags Admin
// @Accept json
// @Produce json
// @Param q query string true "Search keyword (matches title or description)"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "List of matching tasks with count"
// @Failure 400 {object} map[string]string "Search query cannot be empty"
// @Failure 401 {object} map[string]string "Unauthorized (missing or invalid token)"
// @Failure 403 {object} map[string]string "Forbidden (admin access required)"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/searchTask [get]
func SearchTasks(c *gin.Context) {
	searchQuery := strings.TrimSpace(c.Query("q"))

	if searchQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search query cannot be empty"})
		return
	}

	role := c.MustGet("role").(string)
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "only admin can search tasks"})
		return
	}

	var tasks []models.Task
	if err := config.DB.
		Where("title LIKE ? OR description LIKE ?",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%").
		Find(&tasks).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(tasks),
		"tasks": tasks,
	})
}
