package main

import (
	"fmt"
)

func main() {
	app := App{}

	// initialising env vars
	dbUser := loadEnvVar("DB_USER")
	dbPassword := loadEnvVar("DB_PASSWORD")
	dbName := loadEnvVar("DB_NAME")
	appPort := loadEnvVar("APP_PORT")
	dbPort := loadEnvVar("DB_PORT")

	app.Initialise(dbUser, dbPassword, dbPort, dbName)
	app.Run(fmt.Sprintf("localhost:%s", appPort))
}
