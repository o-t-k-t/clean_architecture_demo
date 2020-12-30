module github.com/TechDepa/c_tool/usecase

go 1.11

replace github.com/TechDepa/c_tool/domain/model => ../domain/model

replace github.com/TechDepa/c_tool/adapters/gateways => ./src/adapters/gateways

require (
	github.com/TechDepa/c_tool/domain/model v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/lib/pq v1.8.0
	github.com/pkg/errors v0.9.1
	gopkg.in/gorp.v1 v1.7.2
)
