package controllers

import (
	"net/http"
	"standard_app/models"
	"standard_app/validations"
	"github.com/gin-gonic/gin"
)

type User struct {
}

// @Summary Create a new user
// @Description Create a new user with the given information
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body User true "User data"
// @Success 201 {object} models.User
// @Router /users [post]
func (u *User) Create(c *gin.Context) {
	var user models.User
	var validateUser validations.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate user input and collect all errors
	errors := validateUser.ValidateUser(user)
	// If there are validation errors, return them all at once
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	// Save user
	createdUser, err := user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// @Summary Get all users
// @Description Get a list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 400 {object} map[string]interface{}
// @Router /users [get]
func (u *User) List(c *gin.Context) {
	var user models.User
	userList, err := user.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userList})
}
