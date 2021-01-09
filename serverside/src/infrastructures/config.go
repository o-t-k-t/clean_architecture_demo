package infrastructures

import (
	"fmt"
	"os"
)

func Config(name string) string {
	return os.Getenv(name)
}

func DBDataSource() string {
	var host, port, user, password, dbname, sslmode string
	switch env() {
	case "prod":
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname = os.Getenv("DB_NAME")
		sslmode = "enable"
	case "test":
		host = "localhost"
		port = "5432"
		user = "postgres"
		password = "postgres"
		dbname = "c_tool_test"
		sslmode = "disable"
	default:
		host = "localhost"
		port = "5432"
		user = "postgres"
		password = "postgres"
		dbname = "c_tool_development"
		sslmode = "disable"
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
}

func env() string {
	return os.Getenv("C_TOOL_ENVIRONMENT")
}
