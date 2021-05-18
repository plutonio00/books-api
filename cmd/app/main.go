package main

import (
	"github.com/plutonio00/books-api/internal/app"
)

const configPath = "config/config"

func main() {
	app.Run(configPath)
}
