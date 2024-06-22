package adapters

import (
	"context"
	"io"
	"log"

	"github.com/gocraft/dbr/v2"
	"github.com/irede-interview/cinema-api/internal/config"
	"github.com/irede-interview/cinema-api/internal/database"
	"github.com/irede-interview/cinema-api/internal/repositories"
	"github.com/irede-interview/cinema-api/pkg/logger"
)

func New(conf config.Config) *Adapters {
	a := &Adapters{}

	a.Config = conf

	l := logger.New()
	a.Logger = l

	session := database.Connect(conf.DSN)
	a.DB = session
	a.closers = append(a.closers, session)

	a.Repositories = repositories.New(a.DB)

	a.closers = append(a.closers, l)

	return a
}

type Adapters struct {
	DB      *dbr.Session
	Config  config.Config
	Logger  logger.Provider
	closers []io.Closer

	Repositories repositories.RepoProvider
}

func (a *Adapters) Shutdown(ctx context.Context) {
	for _, c := range a.closers {
		select {
		case <-ctx.Done():
			return
		default:
			err := c.Close()
			if err != nil {
				log.Printf("Error closing component: %v", err)
			}
		}
	}
}
