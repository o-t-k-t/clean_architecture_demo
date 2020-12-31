package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/TechDepa/c_tool/adapters/controllers"
)

type User struct {
	Id int `db:"id"`
}

func main() {
	uc := controllers.NewUsersContorller()

	r := gin.Default()
	r.GET("/users", uc.ShowAll)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ping(gc *gin.Context) {
	fmt.Println("pong!")
}
