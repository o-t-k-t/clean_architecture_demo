package controllers

import (
	"github.com/TechDepa/c_tool/adapters/gateways"
	"github.com/TechDepa/c_tool/infrastructures"
	"github.com/TechDepa/c_tool/usecase"
	"github.com/pkg/errors"
)

type AdminUsersContorller struct{}

func NewAdminUsersController() AdminUsersContorller {
	return AdminUsersContorller{}
}

func (AdminUsersContorller) ShowAll(c Request) {
	infrastructures.WithDatabase(
		func(db infrastructures.Dababase) error {
			r := gateways.NewAdminUsersRepository(db, nil)

			sc, users, err := usecase.ShowAllAdminUsers(r)
			if err != nil {
				c.RenderAbortWithError(sc, err)
				return errors.WithStack(err)
			}
			c.RenderJSON(sc, users)
			return nil
		},
	)
}

func (AdminUsersContorller) Create(c Request) {
	var u usecase.CreateUserInput
	if err := c.BindJSON(&u); err != nil {
		c.RenderAbortWithError(415, err)
		return
	}

	infrastructures.WithDatabaseAndTransaction(
		func(db infrastructures.Dababase, tx infrastructures.Transaction) error {
			r := gateways.NewAdminUsersRepository(db, tx)

			sc, u, err := usecase.CreateUser(u, r)
			if err != nil {
				c.RenderAbortWithError(sc, err)
				return errors.WithStack(err)
			}

			c.RenderJSON(sc, u)
			return nil
		},
	)
}

func (AdminUsersContorller) Update() {

}

func (AdminUsersContorller) Delete() {

}
