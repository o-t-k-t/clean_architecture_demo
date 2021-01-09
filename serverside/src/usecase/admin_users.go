package usecase

import (
	"github.com/TechDepa/c_tool/domain/model"
	"github.com/pkg/errors"
)

type AdminUsersRepository interface {
	FindAll() (model.UserList, error)
	Create(u model.AdminUser) error
}

func ShowAllAdminUsers(r AdminUsersRepository) (model.UserList, error) {
	ul, err := r.FindAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ul, nil
}

func CreateUser(u model.AdminUser, r AdminUsersRepository) (model.AdminUser, error) {
	err := r.Create(u)
	if err != nil {
		return u, errors.WithStack(err)
	}

	return u, nil
}
