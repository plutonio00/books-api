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
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/plutonio00/books-api/pkg/token"
)

func Run(configPath string) {
	conf, err := config.Init(configPath)

	if err != nil {
		logger.Error(err)
		return
	}

	db, err := sql.Open("mysql", conf.Database.MySQL.URI)

	if err != nil {
		logger.Error(err)
		return
	}

	if err := db.Ping(); err != nil {
		logger.Error(err)
		return
	}

	repos := repository.NewRepositories(db)

	tokenManager, err := token.NewTokenManager(conf.Token.JWT.SigningKey)
	if err != nil {
		logger.Error(err)
		return
	}

	services := service.NewServices(
		service.Deps{
			Repos:        repos,
			TokenManager: tokenManager,
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
