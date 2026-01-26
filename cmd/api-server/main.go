package main

import (
	"fmt"
	"my-inventory/internal/app"
	"my-inventory/utils"
)

func main() {
	app := app.App{}
	cfg := utils.LoadEnvConfig()

	app.Initialise(
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	app.Run(fmt.Sprintf(":%s", cfg.AppPort))
}
