package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

type Base struct {
	Id        int64     `db:"id" json:"id, primarykey"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

func (i *Base) PreInsert(s gorp.SqlExecutor) error {
	i.CreatedAt = time.Now()
	i.UpdatedAt = i.CreatedAt
	return nil
}

func (i *Base) PreUpdate(s gorp.SqlExecutor) error {
	i.UpdatedAt = time.Now()
	return nil
}

type BaseUser struct {
	Base
	Email string `db:"email" json:"email"`
	Name  string `db:"name" json:"name"`
}

type AdminUser struct {
	BaseUser
	AdminUserPropaty
}

type AdminUserPropaty struct{}

type UserList []AdminUser
