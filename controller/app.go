package controller

import (
	"online-store-golang/configuration"
	repositoryimplementation "online-store-golang/repository/repository_implementation"
	authservice "online-store-golang/service/auth_service"
	serviceimplementation "online-store-golang/service/service_implementation"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	config := configuration.New()
	db := configuration.NewDatabase(config)

	userRepo := repositoryimplementation.NewUserRepository(db)
	userService := serviceimplementation.NewUserService(userRepo)
	userController := NewUserController(userService)

	productRepo := repositoryimplementation.NewProductRepository(db)
	productService := serviceimplementation.NewProductService(productRepo)
	productController := NewProductController(productService)

	cartRepo := repositoryimplementation.NewCartRepository(db)
	cartService := serviceimplementation.NewCartService(cartRepo, productRepo)
	cartController := NewCartController(cartService)

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
		productRoutes.POST("/", productController.Create)
	}
	cartRoutes := apiV1.Group("/cart")
	{
		cartRoutes.POST("/", authService.Authentication(), cartController.AddToCart)
		cartRoutes.GET("/", authService.Authentication(), cartController.GetUserCart)
		cartRoutes.DELETE("/:productId", authService.Authentication(), cartController.DeleteItem)
	}

	app.Run(":8080")
}
