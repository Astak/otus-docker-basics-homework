package viewmodels

import "github.com/Astak/otus-docker-basics-homework/web-service-gin/data"

type UserRequest struct {
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UserResponse struct {
	ID        int64  `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func (response *UserResponse) MapFromModel(model data.User) {
	response.ID = model.ID
	response.UserName = model.UserName
	response.FirstName = model.FirstName
	response.LastName = model.LastName
	response.Email = model.Email
	response.Phone = model.Phone
}
