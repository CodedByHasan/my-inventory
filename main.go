package main

import (
	"fmt"
)

func main() {
	app := App{}

	db_user := loadEnvVar("DB_USER")
	db_password := loadEnvVar("DB_PASSWORD")
	db := loadEnvVar("DB")
	app_port := loadEnvVar("APP_PORT")

	app.Initialise(db_user, db_password, db)
	app.Run(fmt.Sprintf("localhost:%s", app_port))
}
