package controllers

import (
	"github.com/TechDepa/c_tool/adapters/gateways"
	"github.com/TechDepa/c_tool/usecase"
	"github.com/gin-gonic/gin"
)

type UsersContorller struct{}

func NewUsersContorller() UsersContorller {
	return UsersContorller{}
}

func (UsersContorller) ShowAll(c *gin.Context) {
	r := gateways.NewUsersRepository()
	defer r.Close()

	users, err := usecase.ShowAllUsers(r)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, users)
}

func (UsersContorller) Create() {

}

func (UsersContorller) Update() {

}

func (UsersContorller) Delete() {

}
