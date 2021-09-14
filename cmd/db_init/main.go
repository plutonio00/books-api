package main

import (
    "github.com/plutonio00/books-api/internal/config"
    "os/exec"
    "bytes"
    "fmt"
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
        "up")

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
        conf.Database.MySQL.Host)

    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
     err := cmd.Run()

     if err != nil {
         //panic(err)
         fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
         return
     }
}

//goose -dir internal/migration mysql "books_user:books_pass@tcp(mysql-books:3306)/books_db?parseTime=true"  up
//charlatan load ./tests/fixtures -u=books_user -d=books_db -p=books_pass --host=mysql-books