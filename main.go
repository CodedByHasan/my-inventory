package main

import (
	"fmt"
)

func main() {
	app := App{}

	// initialising env vars
	dbUser := loadEnvVar("DB_USER")
	dbPassword := loadEnvVar("DB_PASSWORD")
	db := loadEnvVar("DB")
	appPort := loadEnvVar("APP_PORT")
	dbPort := loadEnvVar("DB_PORT")

	app.Initialise(dbUser, dbPassword, dbPort, db)
	app.Run(fmt.Sprintf("localhost:%s", appPort))
}
