package main

import (
	"github.com/plutonio00/books-api/internal/config"
	"os/exec"
)

func main() {
	const configPath = "config/config"

	conf, err := config.Init(configPath)

	if err != nil {
		panic(err)
	}

	migrateInit(conf)
	fixturesInit(conf)
}

func migrateInit(conf *config.Config) {
	cmd := exec.Command(
		"goose",
		"-dir",
		"internal/migration",
		"mysql",
		conf.Database.MySQL.DSN,
		"up",
	)

	err := cmd.Run()

	if err != nil {
		panic(err)
		return
	}
}

func fixturesInit(conf *config.Config) {
	cmd := exec.Command(
		"charlatan",
		"load",
		"./tests/fixtures",
		"-u",
		conf.Database.MySQL.User,
		"-d",
		conf.Database.MySQL.DBName,
		"-p",
		conf.Database.MySQL.Password,
		"--host",
		conf.Database.MySQL.Host,
    )

	err := cmd.Run()

	if err != nil {
		panic(err)
		return
	}
}
