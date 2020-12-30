package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/TechDepa/c_tool/adapters/controllers"
)

type User struct {
	Id int `db:"id"`
}

func main() {
	// controllers.Load()

	uc := controllers.NewUsersContorller()

	r := gin.Default()
	r.GET("/users", uc.ShowAll)

	r.GET("/ping", ping)
	r.POST("/_ah/mail/:ToAddress", incomingMail)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ping(gc *gin.Context) {
	fmt.Println("AAAAAAAAAAAAAAAAA")
}

func incomingMail(gc *gin.Context) {
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "test@example.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
