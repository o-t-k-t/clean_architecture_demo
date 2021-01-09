package gateways

import (
	"github.com/pkg/errors"

	"github.com/TechDepa/c_tool/domain/model"
)

type AdminUsersRepository struct {
	db database
	tx transaction
}

// NewAdminUsersRepository AdminUsersRepositoryインスタンスを作成
func NewAdminUsersRepository(database database, transaction transaction) AdminUsersRepository {
	return AdminUsersRepository{
		db: database,
		tx: transaction,
	}
}

func (r AdminUsersRepository) FindAll() (model.UserList, error) {
	q := `
	select * from users
	join offices on users.office_id = offices.id
	join corporations on offices.corporation_id = corporations.id
	order by users.id
	`

	var users []model.AdminUser
	_, err := r.db.Select(&users, q)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return model.UserList(users), nil
}

func (r AdminUsersRepository) Create(u model.AdminUser) error {
	err := r.tx.Insert(u)
	if err != nil {
		return errors.WithMessagef(err, "users挿入失敗", err)
	}
	return nil
}
