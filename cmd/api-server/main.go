package main

import (
	"fmt"
	"my-inventory/internal/app"
	"my-inventory/utils"
)

func main() {
	app := app.App{}
	app.Initialise(utils.LoadEnvConfig())
	app.Run(fmt.Sprintf(":%s", utils.LoadEnvConfig().AppPort))
}
