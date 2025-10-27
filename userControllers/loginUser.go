package controllers

import (
	"net/http"
	config "todo-manager/Config"
	"todo-manager/models"
	"todo-manager/utils"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginUser godoc
// @Summary Login user
// @Description Authenticate user with username and password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param credentials body LoginRequest true "User login credentials"
// @Success 200 {object} map[string]string "Login successful"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 401 {object} map[string]string "Invalid username or password"
// @Router /login [post]
func LoginUser(c *gin.Context) {
	var req LoginRequest

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by username
	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.UserID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// For now, just return success (later we’ll add JWT token here)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    token,
	})
}
