package usecase

import (
	"github.com/TechDepa/c_tool/domain/model"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
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
	Password        model.Password `json:"password" validate:"required,min=6,max=24"`
	PasswordConfirm model.Password `json:"passwordConfirm" validate:"required,min=6,max=24"`
}

func (in CreateUserInput) Validate() error {
	err := validator.New().Struct(in)
	if err != nil {
		return errors.WithStack(err)
	}

	if in.Password != in.PasswordConfirm {
		return errors.New("確認パスワード不一致")
	}

	return nil
}

func CreateUser(in CreateUserInput, r AdminUsersRepository) (model.AdminUser, error) {
	if err := in.Validate(); err != nil {
		return model.AdminUser{}, errors.WithStack(err)
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
