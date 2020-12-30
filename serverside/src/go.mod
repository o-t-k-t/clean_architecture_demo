module github.com/TechDepa/c_tool

go 1.11

replace github.com/TechDepa/c_tool/usecase => ./usecase

replace github.com/TechDepa/c_tool/domain/model => ./domain/model

replace github.com/TechDepa/c_tool/adapters/gateways => ./adapters/gateways

replace github.com/TechDepa/c_tool/adapters/controllers => ./adapters/controllers

replace github.com/TechDepa/c_tool/infrastructures => ./infrastructures

require (
	github.com/TechDepa/c_tool/adapters/controllers v0.0.0-00010101000000-000000000000
	github.com/TechDepa/c_tool/adapters/gateways v0.0.0-00010101000000-000000000000
	github.com/TechDepa/c_tool/domain/model v0.0.0-00010101000000-000000000000
	github.com/TechDepa/c_tool/infrastructures v0.0.0-00010101000000-000000000000
	github.com/TechDepa/c_tool/usecase v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/lib/pq v1.8.0
	github.com/sendgrid/rest v2.6.2+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.7.2+incompatible
	google.golang.org/appengine v1.6.6
	gopkg.in/gorp.v1 v1.7.2
)
