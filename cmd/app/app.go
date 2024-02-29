package main

import "tasktracker/internal/app"

const configsDir = "configs/"

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	app.Run(configsDir, envLocal)
}
