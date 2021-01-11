package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/TechDepa/c_tool/adapters/controllers"
)

func setupRouter() *gin.Engine {
	uc := controllers.NewAdminUsersContorller()

	r := gin.Default()
	r.GET("/v1/admin/users", uc.ShowAll)
	r.POST("/v1/admin/users", uc.Create)

	return r
}

func main() {
	err := setupRouter().Run()
	log.Fatal(errors.WithMessagef(err, "ルーティングが終了した"))
}
