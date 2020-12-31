package usecase

import (
	"github.com/TechDepa/c_tool/domain/model"
	"github.com/pkg/errors"
)

type UsersRepository interface {
	FindAll() (model.UserList, error)
}

func ShowAllUsers(r UsersRepository) (model.UserList, error) {
	ul, err := r.FindAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ul, nil
}
