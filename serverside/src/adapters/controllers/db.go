package controllers

type database interface {
	Begin() (transaction, error)
	Close() error
}

type transaction interface{}
