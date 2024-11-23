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

	productRepo := repositoryimplementation.NewProductRepository(db)
	productService := serviceimplementation.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	app := gin.Default()

	defaultRoutes := app.Group("/api/v1")
	{
		defaultRoutes.POST("/register", userController.Register)
		defaultRoutes.POST("/login", userController.Login)
		defaultRoutes.GET("/products/:categoryId", productController.FindByCategory)
		defaultRoutes.GET("/products", productController.FindAll)

	}

	app.Run(":8080")
}
