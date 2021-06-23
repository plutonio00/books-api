package main

import (
	"github.com/plutonio00/books-api/internal/app"
)

const configPath = "config/config"

// @title Books & Authors API
// @version 1.0

// @host 127.0.0.1:8080
// @BasePath /api/v1/

// @securityDefinitions.apikey ApiTokenAuth
// @in header
// @name Authorization

func main() {
	app.Run(configPath)
}
