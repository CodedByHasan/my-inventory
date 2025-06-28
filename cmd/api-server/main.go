package main

import (
	"fmt"
	"my-inventory/utils"
	"my-inventory/internal/app"
)

func main() {
	app := app.App{}
	cfg := utils.LoadEnvConfig()

	app.Initialise(cfg.DBUser, cfg.DBPassword, cfg.DBPort, cfg.DBName)
	app.Run(fmt.Sprintf("localhost:%s", cfg.AppPort))
}
