package main

import (
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	healthController := controllers.NewHealthController()
	router.GET("/health", healthController.GetHealth)

	userController := controllers.NewUserController()
	router.GET("/user/:id", userController.GetUserHandler)
	router.POST("/user", userController.CreateUserHandler)
	router.PUT("/user/:id", userController.UpdateUserHandler)
	router.DELETE("/user/:id", userController.DeleteUserHandler)

	router.Run("0.0.0.0:8000")
}
