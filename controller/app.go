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
	configuration.HandleRequiredTables(db)

	redisClient := configuration.NewRedis(config)

	userRepo := repositoryimplementation.NewUserRepository(db)
	userService := serviceimplementation.NewUserService(userRepo)
	userController := NewUserController(userService)

	categoryRepo := repositoryimplementation.NewCategoryRepository(db)
	categoryService := serviceimplementation.NewCategoryService(categoryRepo)
	categoryController := NewCategoryController(categoryService)

	productRepo := repositoryimplementation.NewProductRepository(db)
	productService := serviceimplementation.NewProductService(productRepo, categoryRepo, redisClient)
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
		productRoutes.GET("/category/:categoryId", productController.FindByCategory)
		productRoutes.GET("/", productController.FindAll)
		productRoutes.GET("/:productId", productController.FindById)
		productRoutes.POST("/", productController.Create)
		productRoutes.DELETE("/:productId", productController.Delete)
	}
	cartRoutes := apiV1.Group("/cart")
	{
		cartRoutes.POST("/", authService.Authentication(), cartController.AddToCart)
		cartRoutes.GET("/", authService.Authentication(), cartController.GetUserCart)
		cartRoutes.DELETE("/:productId", authService.Authentication(), cartController.DeleteItem)
	}
	categoryRoutes := apiV1.Group("/category")
	{
		categoryRoutes.POST("/", categoryController.Create)
		categoryRoutes.GET("/", categoryController.FindAll)
		categoryRoutes.DELETE("/:categoryId", categoryController.Delete)
	}

	app.Run(":8080")
}
