module github.com/TechDepa/c_tool/adapters/gateways

go 1.11

replace github.com/TechDepa/c_tool/usecase => ../../usecase

replace github.com/TechDepa/c_tool/infrastructures => ../../infrastructures

replace github.com/TechDepa/c_tool/adapters" => ../

replace github.com/TechDepa/c_tool/domain/model => ../../domain/model

require (
	cloud.google.com/go/bigquery v1.13.0
	cloud.google.com/go/storage v1.10.0
	github.com/TechDepa/c_tool/domain/model v0.0.0-00010101000000-000000000000
	github.com/TechDepa/c_tool/infrastructures v0.0.0-00010101000000-000000000000
	github.com/TechDepa/c_tool/usecase v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/lib/pq v1.8.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	gopkg.in/gorp.v1 v1.7.2
)
