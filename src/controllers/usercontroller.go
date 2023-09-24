package controllers

import (
	"net/http"
	"strconv"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/data"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/viewmodels"
	"github.com/gin-gonic/gin"
)

type userController struct {
	store *data.UserStore
}

func NewUserController() *userController {
	store := data.NewUserStore()
	return &userController{store: store}
}

func (controller *userController) CreateUserHandler(c *gin.Context) {
	var request viewmodels.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errModel := viewmodels.NewBadRequest(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	user := controller.store.CreateUser(request.UserName, request.FirstName, request.LastName, request.Email, request.Phone)
	response := viewmodels.UserResponse{}
	response.MapFromModel(user)
	c.IndentedJSON(http.StatusOK, response)
}

func (uc *userController) GetUserHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		errModel := viewmodels.NewBadRequest(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	user, err := uc.store.GetUser(id)
	if err != nil {
		errModel := viewmodels.NewNotFound(err.Error())
		c.IndentedJSON(http.StatusNotFound, errModel)
		return
	}
	response := viewmodels.UserResponse{}
	response.MapFromModel(user)
	c.IndentedJSON(http.StatusOK, response)
}

func (uc *userController) DeleteUserHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		errModel := viewmodels.NewBadRequest(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	if err = uc.store.DeleteUser(id); err != nil {
		errModel := viewmodels.NewNotFound(err.Error())
		c.IndentedJSON(http.StatusNotFound, errModel)
		return
	}
	c.Status(http.StatusNoContent)
}

func (uc *userController) UpdateUserHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		errModel := viewmodels.NewBadRequest(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	var request viewmodels.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errModel := viewmodels.NewBadRequest(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	user, err := uc.store.UpdateUser(id, request.UserName, request.FirstName, request.LastName, request.Email, request.Phone)
	if err != nil {
		errModel := viewmodels.NewNotFound(err.Error())
		c.IndentedJSON(http.StatusNotFound, errModel)
		return
	}
	response := viewmodels.UserResponse{}
	response.MapFromModel(user)
	c.IndentedJSON(http.StatusOK, response)
}
