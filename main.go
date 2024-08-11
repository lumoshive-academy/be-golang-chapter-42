package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// docs is generated by Swag CLI, you have to import it.
	"golang-chapter-42/controller"
	_ "golang-chapter-42/docs"
)

// @title Example API
// @version 1.0
// @description This is a sample server for a Swagger API.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url https://academy.lumoshive.com/contact-us
// @contact.email lumoshive.academy@gmail.com
// @license.name Lumoshive Academy
// @license.url https://academy.lumoshive.com
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	ctl := controller.NewControllerUser()

	// Swagger endpoint
	r.GET("/users/:id", ctl.GetByID)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
