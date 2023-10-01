package router

import (
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handler.Handler) *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", handler.GetHealth)
	engine.GET("/user/:id", handler.GetUser)
	engine.POST("/user", handler.CreateUser)
	engine.DELETE("/user/:id", handler.DeleteUser)
	engine.PUT("/user/:id", handler.UpdateUser)
	return engine
}
