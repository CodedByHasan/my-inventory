package main

import (
	"fmt"
	"my-inventory/utils"
	"my-inventory/internal/app"
)

func main() {
	app := app.App{}

	// initialising env vars
	dbUser := utils.LoadEnvVar("DB_USER")
	dbPassword := utils.LoadEnvVar("DB_PASSWORD")
	dbName := utils.LoadEnvVar("DB_NAME")
	appPort := utils.LoadEnvVar("APP_PORT")
	dbPort := utils.LoadEnvVar("DB_PORT")

	app.Initialise(dbUser, dbPassword, dbPort, dbName)
	app.Run(fmt.Sprintf("localhost:%s", appPort))
}
