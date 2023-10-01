package handler

import (
	"net/http"
	"strconv"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/dto"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	request := new(dto.UserRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	u := &models.User{
		UserName:  request.UserName,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
	}
	newUser, err := h.UserRepo.CreateUser(u)
	if err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	response := dto.UserResponse{
		ID:        newUser.ID,
		UserName:  newUser.UserName,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Phone:     newUser.Phone,
	}
	c.IndentedJSON(http.StatusCreated, response)
}

func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	user, err := h.UserRepo.GetUser(id)
	if err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	if user == nil {
		errModel := dto.NewNotFoundError("User not found")
		c.IndentedJSON(http.StatusNotFound, errModel)
		return
	}
	response := dto.UserResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	err = h.UserRepo.DeleteUser(id)
	if err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	request := new(dto.UserRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	u := &models.User{
		UserName:  request.UserName,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
	}
	newUser, err := h.UserRepo.UpdateUser(id, u)
	if err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	if newUser == nil {
		errModel := dto.NewNotFoundError("User not found")
		c.IndentedJSON(http.StatusNotFound, errModel)
		return
	}
	response := dto.UserResponse{
		ID:        newUser.ID,
		UserName:  newUser.UserName,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Phone:     newUser.Phone,
	}
	c.IndentedJSON(http.StatusOK, response)
}
