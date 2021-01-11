package model

import (
	"time"

	"gopkg.in/gorp.v1"
)

var overRappingNowTime *time.Time

// NowTime 現在時刻の取得
func NowTime() time.Time {
	if overRappingNowTime != nil {
		return *overRappingNowTime
	}
	return time.Now()
}

// OverrapNowTime テストスタブ用に現在時刻を上書き更新
func OverrapNowTime(year int, month time.Month, day, hour, min, sec, nsec int) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	t := time.Date(year, month, day, hour, min, sec, nsec, jst)
	overRappingNowTime = &t
}

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
