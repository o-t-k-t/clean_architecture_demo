package usecase

import (
	"github.com/TechDepa/c_tool/domain/model"
	"github.com/pkg/errors"
)

type AdminUsersRepository interface {
	FindAll() (model.AdminUserList, error)
	Create(u model.AdminUser) error
}

func ShowAllAdminUsers(r AdminUsersRepository) (model.AdminUserList, error) {
	ul, err := r.FindAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ul, nil
}

type CreateUserInput struct {
	AdminUser       model.AdminUser
	Password        model.Password `json:"password"`
	PasswordConfirm model.Password `json:"passwordConfirm"`
}

func (in CreateUserInput) Valid() bool {
	if in.Password != in.PasswordConfirm {
		return false
	}

	if !in.Password.Valid() {
		return false
	}

	if !in.PasswordConfirm.Valid() {
		return false
	}
	return true
}

var InvalidAdminUserError = errors.New("管理ユーザー入力データ不正")

func CreateUser(in CreateUserInput, r AdminUsersRepository) (model.AdminUser, error) {
	if !in.Valid() {
		return model.AdminUser{}, InvalidAdminUserError
	}

	au := in.AdminUser
	hash, err := in.Password.NewPasswordHash()
	if err != nil {
		return au, errors.WithStack(err)
	}
	au.PasswordHash = hash

	err = r.Create(au)
	if err != nil {
		return au, errors.WithStack(err)
	}

	return au, nil
}
