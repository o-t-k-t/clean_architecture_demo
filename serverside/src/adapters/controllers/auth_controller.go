package controllers

import (
	"github.com/TechDepa/c_tool/adapters/gateways"
	"github.com/TechDepa/c_tool/domain/model"
	"github.com/TechDepa/c_tool/usecase"
	"github.com/pkg/errors"
)

type AuthController struct {
	db gateways.Database
}

func NewAuthController(db gateways.Database) AuthController {
	return AuthController{db}
}

func (c AuthController) Authenticator(
	in usecase.AuthenticatorInput,
	db gateways.Database,
) (model.AdminUser, error) {
	db.BeginConnection()
	defer db.Close()

	// if err := json.NewDecoder(c.Request.Body).Decode(&loginVals); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := req.BindJSON(&in); err != nil {
	// 	return model.AdminUser{}, errors.WithStack(err)
	// }

	r := gateways.NewAdminUsersRepository(c.db)

	_, user, err := usecase.Authenticator(in, r)
	if err != nil {
		return model.AdminUser{}, errors.WithStack(err)
	}

	return user, nil
}
