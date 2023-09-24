package controllers

import (
	"net/http"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/viewmodels"
	"github.com/gin-gonic/gin"
)

type healthController struct{}

func NewHealthController() *healthController {
	return &healthController{}
}

func (controller *healthController) GetHealth(c *gin.Context) {
	health := viewmodels.NewHealthOk()
	c.IndentedJSON(http.StatusOK, health)
}
