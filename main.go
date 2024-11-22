package main

import (
	"online-store-golang/configuration"
	"online-store-golang/controller"
	repositoryimplementation "online-store-golang/repository/repository_implementation"
	serviceimplementation "online-store-golang/service/service_implementation"

	"github.com/gin-gonic/gin"
)

func main() {

	config := configuration.New()
	db := configuration.NewDatabase(config)

	userRepo := repositoryimplementation.NewUserRepository(db)
	userService := serviceimplementation.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	app := gin.Default()

	defaultRoutes := app.Group("/api/v1")
	{
		defaultRoutes.POST("/register", userController.Register)
	}

	app.Run(":8080")
}
