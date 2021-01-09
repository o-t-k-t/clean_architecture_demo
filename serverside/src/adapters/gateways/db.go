package gateways

type database interface {
	Select(i interface{}, query string, args ...interface{}) ([]interface{}, error)
}

type transaction interface {
	Commit() error
	Insert(list ...interface{}) error
}
