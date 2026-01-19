package controllers

import (
	"net/http"
	"time"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

// AdminGetAllTasks godoc
// @Summary      Get all tasks (Admin only)
// @Description  Admin can view all tasks across all users
// @Tags         Admin
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} map[string][]models.Task "List of tasks"
// @Failure      403 {object} map[string]string "Admin access only"
// @Failure      500 {object} map[string]string "Failed to fetch tasks"
// @Router       /admin/tasks [get]
func AdminGetAllTasks(c *gin.Context) {

	// Admin-only check
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "admin access only"})
		return
	}

	var tasks []models.Task

	if err := config.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	// Convert UTC → IST
	loc, _ := time.LoadLocation("Asia/Kolkata")
	for i := range tasks {
		tasks[i].DueDate = tasks[i].DueDate.In(loc)
		tasks[i].CreatedAt = tasks[i].CreatedAt.In(loc)
		tasks[i].UpdatedAt = tasks[i].UpdatedAt.In(loc)
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(tasks),
		"tasks": tasks,
	})
}
