package router

import (
	ginprometheus "github.com/Astak/go-gin-prometheus"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handler.Handler) *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", handler.GetHealth)

	p := ginprometheus.NewPrometheus("gin")
	p.Use(engine)

	engine.GET("/user/:id", handler.GetUser)
	engine.POST("/user", handler.CreateUser)
	engine.DELETE("/user/:id", handler.DeleteUser)
	engine.PUT("/user/:id", handler.UpdateUser)
	return engine
}
