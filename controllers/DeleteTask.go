package controllers

import (
	"net/http"
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by its id
// @Tags tasks
// @Param id path int true "Task id"
// @Success 200 {object} map[string]string "Task deleted successfully"
// @Failure 400 {object} map[string]string "Invalid task id"
// @Failure 404 {object} map[string]string "Task not found"
// @Failure 500 {object} map[string]string "Failed to delete task"
// @Router /deleteTask/{id} [delete]
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := config.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
