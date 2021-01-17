package gateways

import "github.com/TechDepa/c_tool/domain/model"

type Database interface {
	BeginConnection()
	BeginConnectionAndTransaction() error
	CommitOrRollbackAndClose(commit bool)
	Close()
	Select(i interface{}, query string, args ...interface{}) ([]interface{}, error)
	SelectOne(holder interface{}, query string, args ...interface{}) error
	Insert(list ...interface{}) error
}

type timer interface {
	Now() model.AppTime
}
