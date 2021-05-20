package app

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/plutonio00/books-api/internal/config"
	delivery "github.com/plutonio00/books-api/internal/delivery/http"
	"github.com/plutonio00/books-api/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/plutonio00/books-api/internal/repository"
	"github.com/plutonio00/books-api/internal/service"
)

func Run(configPath string) {
	conf, err := config.Init(configPath)

	if err != nil {
		fmt.Print(err)
		return
	}

	db, err := sql.Open("mysql", conf.DatabaseConfig.MySQLConfig.URI)

	if err != nil {
		fmt.Print(err)
		return
	}

	if err := db.Ping(); err != nil {
		fmt.Print(err)
		return
	}

	repos := repository.NewRepositories(db)
	services := service.NewServices(
	    service.Deps {
	        Repos: repos,
	    },
	)

	handler := delivery.NewHandler(services)

	srv := server.NewServer(conf, handler.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {

		}
	}()

	fmt.Print("server run")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
