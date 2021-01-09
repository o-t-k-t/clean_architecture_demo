package gateways

import "github.com/TechDepa/c_tool/domain/model"

type database interface {
	Select(i interface{}, query string, args ...interface{}) ([]interface{}, error)
}

type transaction interface {
	Commit() error
	Insert(list ...interface{}) error
}

type timer interface {
	Now() model.AppTime
}
