package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/irede-interview/cinema-api/cmd/http"
	"github.com/irede-interview/cinema-api/internal/adapters"
	"github.com/irede-interview/cinema-api/internal/config"
	"github.com/irede-interview/cinema-api/internal/handlers/healthy"
	"github.com/irede-interview/cinema-api/internal/handlers/moviehdl"
	"github.com/irede-interview/cinema-api/internal/handlers/sessionhdl"
	"github.com/irede-interview/cinema-api/internal/handlers/threaterhdl"
)

func main() {
	conf := config.New()
	apt := adapters.New(conf)

	app := server.New(apt)

	app.RegisterHandler(
		healthy.NewHandler(apt),
		moviehdl.NewHandler(apt),
		sessionhdl.NewHandler(apt),
		threaterhdl.NewHandler(apt),
	)

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := app.Run()
		if err != nil {
			stopCh <- syscall.SIGTERM
		}
	}()
	<-stopCh

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	apt.Shutdown(ctx)
	app.Shutdown(ctx)

	err := app.Router.Run(app.Server.Addr)
	if err != nil {
		panic(fmt.Sprintf("Unable to listen server: %v", err))
	}
}
