package usecase

import (
	"github.com/TechDepa/c_tool/domain/model"
	"github.com/pkg/errors"
)

type AuthenticatorInput struct {
	EMail    string         `json:"email" validate:"required`
	Password model.Password `json:"password" validate:"required,min=6,max=24"`
}

func Authenticator(
	in AuthenticatorInput,
	r AdminUsersRepository,
) (Status, model.AdminUser, error) {
	u, err := r.FindByEmail(in.EMail)
	if err != nil {
		return 500, model.AdminUser{}, errors.WithStack(err)
	}

	if err := in.Password.MatchesHash(u.Hash); err != nil {
		return 401, model.AdminUser{}, errors.New("認証エラー")
	}

	return 200, u, nil
}
