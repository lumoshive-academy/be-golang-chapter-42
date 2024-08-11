package controller

import (
	"golang-chapter-42/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewControllerUser() *Controller {
	return &Controller{}
}

// Get user all endpoint
// @Summary Get user all
// @Description Get details of all registered users.
// @Tags users
// @Accept  json
// @Produce  json
// @Router  /users [get]
func (c *Controller) GetAll(ctx *gin.Context) {

	// response
	ctx.JSON(200, gin.H{"status_code": 200, "message": "show all data users"})
}

// Get user by ID endpoint
// @Summary Get user by ID
// @Description Get details of a user by their ID.
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id     path     int     true  "User ID"
// @Param Authorization header string true "Bearer token"
// @Router  /users/{id} [get]
func (c *Controller) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	// response

	// Dummy implementation for example
	if id == "1" {
		ctx.JSON(200, gin.H{"status_code": 200, "message": "show data user with id"})
	} else {
		ctx.JSON(404, gin.H{"status_code": 404, "error": "User not found"})
	}
}

// store data user endpoint
// @Summary Add a new user
// @Description Add a new user with the provided JSON data.
// @Tags users
// @Param user body model.User true "User data"
// @Param Authorization header string true "Bearer token"
// @Router /users [post]
func (c *Controller) Store(ctx *gin.Context) {
	var newUser model.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// Dummy implementation
	newUser.Id = 3
	ctx.JSON(201, newUser)
}
