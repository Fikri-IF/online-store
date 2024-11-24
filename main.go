package main

import (
	"online-store-golang/configuration"
	"online-store-golang/controller"
	repositoryimplementation "online-store-golang/repository/repository_implementation"
	authservice "online-store-golang/service/auth_service"
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

	cartRepo := repositoryimplementation.NewCartRepository(db)
	cartService := serviceimplementation.NewCartService(cartRepo)
	cartController := controller.NewCartController(cartService)

	authService := authservice.NewAuthService(userRepo)

	app := gin.Default()
	apiV1 := app.Group("/api/v1")

	userRoutes := apiV1.Group("/user")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
	}
	productRoutes := apiV1.Group("/product")
	{
		productRoutes.GET("/products/:categoryId", productController.FindByCategory)
		productRoutes.GET("/products", productController.FindAll)
	}
	cartRoutes := apiV1.Group("/cart")
	{
		cartRoutes.POST("/", authService.Authentication(), cartController.AddToCart)
	}

	app.Run(":8080")
}
