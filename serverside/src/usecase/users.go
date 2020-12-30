package usecase

import (
	"github.com/TechDepa/c_tool/domain/model"
)

type UsersRepository interface {
}

func ShowAllUsers(r UsersRepository) (model.UserList, error) {
	return model.UserList{
		{ID: 1},
		{ID: 2},
	}, nil
}
