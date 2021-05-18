package app

import (
	"errors"
	"fmt"
	"github.com/plutonio00/books-api/internal/config"
	"github.com/plutonio00/books-api/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	conf, err := config.Init(configPath)

	if err != nil {
		fmt.Print(err)
	}

	srv := server.NewServer(conf)

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {

		}
	}()

	fmt.Print("server run")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
