package controllers

import (
	"github.com/TechDepa/c_tool/adapters/gateways"
	"github.com/TechDepa/c_tool/usecase"
	"github.com/pkg/errors"
)

type database interface {
	BeginConnection()
	BeginConnectionAndTransaction() error
	Close(commit bool) error
}

type AdminUsersController struct {
	db gateways.Database
}

func NewAdminUsersController(db gateways.Database) AdminUsersController {
	return AdminUsersController{db}
}

func (c AdminUsersController) ShowAll(req Request) {
	c.db.BeginConnection()
	defer c.db.Close()

	r := gateways.NewAdminUsersRepository(c.db)

	sc, users, err := usecase.ShowAllAdminUsers(r)
	if err != nil {
		req.RenderAbortWithError(sc, err)
		return
	}

	req.RenderJSON(sc, users)
	return
}

func (c AdminUsersController) Create(req Request) {
	var in usecase.CreateUserInput
	if err := req.BindJSON(&in); err != nil {
		req.RenderAbortWithError(415, errors.WithStack(err))
		return
	}

	r := gateways.NewAdminUsersRepository(c.db)

	commit := false
	if err := c.db.BeginConnectionAndTransaction(); err != nil {
		req.RenderAbortWithError(500, errors.WithStack(err))
		return
	}
	defer c.db.CommitOrRollbackAndClose(commit)

	sc, u, err := usecase.CreateUser(in, r)
	if err != nil {
		req.RenderAbortWithError(sc, err)
		return
	}

	req.RenderJSON(sc, u)
	commit = true

	return
}

func (AdminUsersController) Update() {

}

func (AdminUsersController) Delete() {

}
