package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/irede-interview/cinema-api/internal/adapters"
)

type App struct {
	Adapters *adapters.Adapters
	Server   *http.Server
	Router   *gin.Engine
}

type HandlerRegisterer interface {
	SetUpRoutes(*gin.Engine)
}

func (app *App) RegisterHandler(registers ...HandlerRegisterer) {
	for _, registerer := range registers {
		registerer.SetUpRoutes(app.Router)
	}
}

func (app *App) Run() error {
	return app.Server.ListenAndServe()
}

func (app *App) Shutdown(ctx context.Context) {
	err := app.Server.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}

func New(apt *adapters.Adapters) *App {
	router := gin.New()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	app := &App{
		Adapters: apt,
		Router:   router,
		Server: &http.Server{
			Addr:              fmt.Sprintf("127.0.0.1:%s", apt.Config.Port),
			Handler:           router,
			ReadTimeout:       time.Second * 15,
			ReadHeaderTimeout: time.Second * 15,
			WriteTimeout:      time.Second * 15,
			IdleTimeout:       time.Second * 30,
		},
	}

	return app
}
