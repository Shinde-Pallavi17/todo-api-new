package controllers

import (
	"net/http"
	config "todo-manager/Config"
	"todo-manager/models"
	"todo-manager/utils"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=4"`
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Creates a new user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body RegisterRequest true "User registration data"
// @Success 201 {object} map[string]string "User registered successfully"
// @Failure 400 {object} map[string]string "Invalid input or username already exists"
// @Failure 500 {object} map[string]string "Server error (failed to hash password)"
// @Router /register [post]
func RegisterUser(c *gin.Context) {
	var req RegisterRequest

	// Bind JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create new user
	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
	}

	// Save to DB
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
