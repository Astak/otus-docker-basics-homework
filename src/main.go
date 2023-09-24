package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type health struct {
	Status string `json:"status"`
}

type user struct {
	ID        int64  `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

const (
	StatusOk = "OK"
)

var users []user

func getHealth(c *gin.Context) {
	response := health{
		Status: StatusOk,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func postUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		response := error{
			Code:    400,
			Message: err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response := error{
			Code:    400,
			Message: err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}
	for _, user := range users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	response := error{
		Code:    404,
		Message: "user not found",
	}
	c.IndentedJSON(http.StatusNotFound, response)
}

func deleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response := error{
			Code:    400,
			Message: err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}
	for i, user := range users {
		if user.ID == id {
			users[i] = users[len(users)-1]
			users = users[:len(users)-1]
			c.Status(http.StatusNoContent)
			return
		}
	}
	response := error{
		Code:    404,
		Message: "user not found",
	}
	c.IndentedJSON(http.StatusNotFound, response)
}

func updateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response := error{
			Code:    400,
			Message: err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}

	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		response := error{
			Code:    400,
			Message: err.Error(),
		}
		c.IndentedJSON(http.StatusBadRequest, response)
		return
	}
	for i, user := range users {
		if user.ID == id {
			users[i] = newUser
			c.IndentedJSON(http.StatusOK, newUser)
			return
		}
	}
	response := error{
		Code:    404,
		Message: "user not found",
	}
	c.IndentedJSON(http.StatusNotFound, response)
}

func main() {
	router := gin.Default()
	router.GET("/health", getHealth)
	router.GET("/user/:id", getUserByID)
	router.POST("/user", postUser)
	router.PUT("/user/:id", updateUser)
	router.DELETE("/user/:id", deleteUser)
	router.Run("0.0.0.0:8000")
}
