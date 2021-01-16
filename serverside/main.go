package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/TechDepa/c_tool/adapters/controllers"
	"github.com/TechDepa/c_tool/infrastructures"
)

func setupRouter() *gin.Engine {
	uc := controllers.NewAdminUsersController()

	r := gin.Default()
	r.GET("/v1/admin/users", func(c *gin.Context) {
		uc.ShowAll(infrastructures.Request{c})
	})
	r.POST("/v1/admin/users", func(c *gin.Context) {
		uc.Create(infrastructures.Request{c})
	})

	return r
}

func main() {
	err := setupRouter().Run()
	log.Fatal(errors.WithMessagef(err, "ルーティングが終了した"))
}
