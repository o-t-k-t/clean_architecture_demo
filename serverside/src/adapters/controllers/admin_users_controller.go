package controllers

import (
	"github.com/TechDepa/c_tool/adapters/gateways"
	"github.com/TechDepa/c_tool/domain/model"
	"github.com/TechDepa/c_tool/infrastructures"
	"github.com/TechDepa/c_tool/usecase"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type AdminUsersContorller struct{}

func NewAdminUsersContorller() AdminUsersContorller {
	return AdminUsersContorller{}
}

func (AdminUsersContorller) ShowAll(c *gin.Context) {
	infrastructures.WithDatabase(
		func(db infrastructures.Dababase) error {
			r := gateways.NewAdminUsersRepository(
				db,
				nil,
				infrastructures.NewTimer(),
			)

			users, err := usecase.ShowAllAdminUsers(r)
			if err != nil {
				c.AbortWithError(500, err)
				return errors.WithStack(err)
			}
			c.JSON(200, users)
			return nil
		},
	)
}

func (AdminUsersContorller) Create(c *gin.Context) {
	var u model.AdminUser
	if err := c.BindJSON(&u); err != nil {
		return
	}

	infrastructures.WithDatabaseAndTransaction(
		func(db infrastructures.Dababase, tx infrastructures.Transaction) error {
			r := gateways.NewAdminUsersRepository(
				db,
				tx,
				infrastructures.NewTimer(),
			)

			u, err := usecase.CreateUser(u, r)
			if err != nil {
				c.AbortWithError(500, err)
				return errors.WithStack(err)
			}
			c.JSON(200, u)
			return nil
		},
	)
}

func (AdminUsersContorller) Update() {

}

func (AdminUsersContorller) Delete() {

}
