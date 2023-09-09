package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type health struct {
	Status string `json:"status"`
}

const (
	StatusOk = "OK"
)

func getHealth(c *gin.Context) {
	response := health{
		Status: StatusOk,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.GET("/health", getHealth)
	router.Run("localhost:8000")
}
