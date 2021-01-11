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
	i.CreatedAt = NowTime()
	i.UpdatedAt = i.CreatedAt
	return nil
}

func (i *Base) PreUpdate(s gorp.SqlExecutor) error {
	i.UpdatedAt = NowTime()
	return nil
}
