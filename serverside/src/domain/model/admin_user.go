package model

type AdminUser struct {
	ID                int64 `db:"id" json:"id"`
	ContranctOfficeID int64 `db:"contranct_office_id" json:"-"`
}

type UserList []AdminUser
