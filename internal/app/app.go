package app

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/plutonio00/books-api/internal/config"
	delivery "github.com/plutonio00/books-api/internal/delivery/http"
	"github.com/plutonio00/books-api/internal/repository"
	"github.com/plutonio00/books-api/internal/server"
	"github.com/plutonio00/books-api/internal/service"
	"github.com/plutonio00/books-api/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	conf, err := config.Init(configPath)

	if err != nil {
		logger.Error(err)
		return
	}

	db, err := sql.Open("mysql", conf.DatabaseConfig.MySQLConfig.URI)

	if err != nil {
		logger.Error(err)
		return
	}

	if err := db.Ping(); err != nil {
		logger.Error(err)
		return
	}

	repos := repository.NewRepositories(db)
	services := service.NewServices(
		service.Deps{
			Repos: repos,
		},
	)

	handler := delivery.NewHandler(services)

	srv := server.NewServer(conf, handler.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {

		}
	}()

	logger.Info("server run")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
