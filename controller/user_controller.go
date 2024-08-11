package controller

import "github.com/gin-gonic/gin"

type Controller struct {
}

func NewControllerUser() *Controller {
	return &Controller{}
}

// Get user by ID endpoint
// @Summary Get user by ID
// @Description Get details of a user by their ID.
// @Tags users
// @Accept  json
// @Produce  json
// @Router  /users/{id} [get]
func (c *Controller) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	// Dummy implementation for example
	if id == "1" {
		ctx.JSON(200, gin.H{"id": id, "name": "lumoshive-academy"})
	} else {
		ctx.JSON(404, gin.H{"error": "User not found"})
	}
}
