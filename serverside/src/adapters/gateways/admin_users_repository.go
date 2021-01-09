package gateways

import (
	"github.com/pkg/errors"

	"github.com/TechDepa/c_tool/domain/model"
)

type AdminUsersRepository struct {
	db    database
	tx    transaction
	timer timer
}

type AdminUserRecord struct {
	model.Base
	BaseUserId int64 `db:"base_user_id"`
	model.AdminUserPropaty
}

// NewAdminUsersRepository AdminUsersRepositoryインスタンスを作成
func NewAdminUsersRepository(database database, transaction transaction, timer timer) AdminUsersRepository {
	return AdminUsersRepository{
		db:    database,
		tx:    transaction,
		timer: timer,
	}
}

func (r AdminUsersRepository) FindAll() (model.UserList, error) {
	q := `
	select
      admin_users.id,
      admin_users.created_at,
      admin_users.updated_at,
      base_users.email,
      base_users.name
	from admin_users
	join base_users on admin_users.base_user_id = base_users.id
	order by admin_users.id
	`

	users := []model.AdminUser{}
	_, err := r.db.Select(&users, q)
	if err != nil {
		return nil, errors.WithMessagef(err, "admin_users取得失敗")
	}

	return model.UserList(users), nil
}

func (r AdminUsersRepository) Create(au model.AdminUser) error {
	bu := au.BaseUser
	if err := r.tx.Insert(&bu); err != nil {
		return errors.WithMessagef(err, "base_users登録失敗")
	}

	aur := AdminUserRecord{BaseUserId: bu.Id}
	if err := r.tx.Insert(&aur); err != nil {
		return errors.WithMessagef(err, "admin_users登録失敗")
	}

	return nil
}
