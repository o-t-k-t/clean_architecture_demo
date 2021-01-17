package model

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Password string

func (v Password) NewPasswordHash() (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(v), 10)
	if err != nil {
		return "", errors.WithMessagef(err, "パスワードハッシュ生成失敗")
	}

	return string(hashed), nil
}

func (v Password) MatchesHash(hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(v))
	if err != nil {
		return errors.WithMessagef(err, "パスワードハッシュ生成失敗")
	}

	return nil
}

type BaseUser struct {
	Base
	Email string `db:"email" json:"email"`
	Name  string `db:"name" json:"name"`
	Hash  string `db:"password_hash" json:"-"`
}

type AdminUser struct {
	BaseUser
	AdminUserPropaty
}

type AdminUserPropaty struct{}

type AdminUserList []AdminUser
