package gateways

import (
	"github.com/pkg/errors"
	"gopkg.in/gorp.v1"

	"github.com/TechDepa/c_tool/domain/model"
	"github.com/TechDepa/c_tool/infrastructures"
)

type UsersRepository struct {
	dbmap *gorp.DbMap
}

func NewUsersRepository() UsersRepository {
	return UsersRepository{
		dbmap: infrastructures.InitDB(),
	}
}

func (r UsersRepository) Close() {
	r.dbmap.Db.Close()
}

func (r UsersRepository) FindAll() (model.UserList, error) {
	q := `
	select * from users
	join offices on users.office_id = offices.id
	join corporations on offices.corporation_id = corporations.id
	order by users.id
	`

	var users []model.User
	_, err := r.dbmap.Select(&users, q)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return model.UserList(users), nil
}

func (r UsersRepository) Create() (model.UserList, error) {
	findAllQuery := `
	select * from users
	join offices on users.office_id = offices.id
	join corporations on offices.corporation_id = corporations.id
	order by users.id
	`

	var users []model.User
	_, err := r.dbmap.Select(&users, findAllQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return model.UserList(users), nil
}
