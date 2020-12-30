module github.com/TechDepa/c_tool/adapters/controllers

go 1.11

replace github.com/TechDepa/c_tool/usecase => ../../usecase

replace github.com/TechDepa/c_tool/domain/model => ../../domain/model

replace github.com/TechDepa/c_tool/adapters/gateways => ../gateways

require (
	github.com/TechDepa/c_tool/usecase v0.0.0-00010101000000-000000000000
	github.com/TechDepa/c_tool/domain/model v0.0.0-00010101000000-000000000000
	github.com/TechDepa/c_tool/adapters/gateways v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/lib/pq v1.8.0
	gopkg.in/gorp.v1 v1.7.2
)
