package controllers

import (
	"net/http"
	"time"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

// GetTasksByGroup godoc
// @Summary Get tasks by group
// @Description Get tasks filtered by specific groups like personal, education, office, etc.
// @Tags tasks
// @Security BearerAuth
// @Produce json
// @Param group path string true "Task Group" Enums(personal,office,shopping,family,friends,education,health,travel,food)
// @Success 200 {array} models.Task
// @Router /tasks/group/{group} [get]
func GetTasksByGroup(c *gin.Context) {
	group := c.Param("group")
	userID := c.GetUint("userID")

	//only allow predefined group categories
	validCategories := map[string]bool{
		"personal":  true,
		"office":    true,
		"shopping":  true,
		"family":    true,
		"friends":   true,
		"education": true,
		"health":    true,
		"travel":    true,
		"food":      true,
	}

	if !validCategories[group] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	var tasks []models.Task

	//Fetch tasks by group
	if err := config.DB.Where("user_id = ? AND task_group = ?", userID, group).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	//Convert dates to IST
	loc, _ := time.LoadLocation("Asia/Kolkata")
	for i := range tasks {
		tasks[i].DueDate = tasks[i].DueDate.In(loc)
		tasks[i].CreatedAt = tasks[i].CreatedAt.In(loc)
		tasks[i].UpdatedAt = tasks[i].UpdatedAt.In(loc)
	}

	c.JSON(http.StatusOK, gin.H{
		"group": group,
		"count": len(tasks),
		"tasks": tasks,
	})
}
