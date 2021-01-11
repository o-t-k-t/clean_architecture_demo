package model

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Password string

func (v Password) Valid() bool {
	if len(v) < 6 || 24 <= len(v) {
		return false
	} else {
		return true
	}
}

func (v Password) NewPasswordHash() (PasswordHash, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(v), 10)
	if err != nil {
		return "", errors.WithMessagef(err, "パスワードハッシュ生成失敗")
	}

	return PasswordHash(hashed), nil
}

func (v Password) MatchesPasswordHash(hash PasswordHash) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(v))
	if err != nil {
		return errors.WithMessagef(err, "パスワードハッシュ生成失敗")
	}

	return nil
}

type PasswordHash string
type BaseUser struct {
	Base
	Email        string       `db:"email" json:"email"`
	Name         string       `db:"name" json:"name"`
	PasswordHash PasswordHash `db:"password_hash" json:"-"`
}

type AdminUser struct {
	BaseUser
	AdminUserPropaty
}

type AdminUserPropaty struct{}

type AdminUserList []AdminUser
