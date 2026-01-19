package controllers

import (
	config "todo-manager/Config"
	"todo-manager/models"

	"github.com/gin-gonic/gin"
)

// AdminGetAllUsers godoc
// @Summary      Get all users (Admin only)
// @Description  Admin can view the list of all registered users
// @Tags         Admin
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} map[string][]models.User "List of users"
// @Failure      403 {object} map[string]string "Admin access only"
// @Failure      500 {object} map[string]string "Failed to fetch users"
// @Router       /admin/users [get]
func AdminGetAllUsers(c *gin.Context) {

	// Role check
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(403, gin.H{"error": "Admin access only"})
		return
	}

	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}
