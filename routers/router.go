package routers

import (
	"standard_app/controllers"
	"standard_app/middlewares"
	"github.com/gin-gonic/gin"
	"standard_app/swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	swagger.RegisterSwagger(router)

	authController := &controllers.Auth{}
	userController := &controllers.User{}
	productController := &controllers.Product{}
	// Public routes (No authentication required)
	router.POST("/auth/login", authController.Login)

	// Group routes that require authentication
	// users
	users := router.Group("/users")
	users.Use(middlewares.AuthMiddleware())
	{
		users.GET("/", userController.List)
		users.POST("/", userController.Create)
	}

	// products
	products := router.Group("/products")
	products.Use(middlewares.AuthMiddleware())
	{
		products.GET("/", productController.List)
		products.GET("/:id", productController.FindById)
		products.POST("/", productController.Create)
	}

	return router
}
