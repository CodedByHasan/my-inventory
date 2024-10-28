package main

import (
	"fmt"
)

func main() {
	app := App{}
	app_port := loadEnvVar("APP_PORT")

	app.Initialise()
	app.Run(fmt.Sprintf("localhost:%s", app_port))
}
