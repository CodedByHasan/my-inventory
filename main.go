package main

import (
	"fmt"
)

func main() {
	app := App{}

	dbUser := loadEnvVar("DB_USER")
	dbPassword := loadEnvVar("DB_PASSWORD")
	db := loadEnvVar("DB")
	appPort := loadEnvVar("APP_PORT")

	app.Initialise(dbUser, dbPassword, db)
	app.Run(fmt.Sprintf("localhost:%s", appPort))
}
