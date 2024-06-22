package repositories

import (
	"github.com/gocraft/dbr/v2"
	"github.com/irede-interview/cinema-api/internal/core/ports"
	"github.com/irede-interview/cinema-api/internal/repositories/movierepo"
	"github.com/irede-interview/cinema-api/internal/repositories/sessionrepo"
	"github.com/irede-interview/cinema-api/internal/repositories/threaterrepo"
)

type RepoProvider interface {
	Movie() ports.MovieRepository
	Session() ports.SessionRepository
	Threater() ports.ThreaterRepository
}

type Provider struct {
	MovieRepository    ports.MovieRepository
	SessionRepository  ports.SessionRepository
	ThreaterRepository ports.ThreaterRepository
}

func (p Provider) Movie() ports.MovieRepository       { return p.MovieRepository }
func (p Provider) Session() ports.SessionRepository   { return p.SessionRepository }
func (p Provider) Threater() ports.ThreaterRepository { return p.ThreaterRepository }

func New(dbrConn *dbr.Session) *Provider {
	return &Provider{
		MovieRepository:    movierepo.New(dbrConn),
		SessionRepository:  sessionrepo.New(dbrConn),
		ThreaterRepository: threaterrepo.New(dbrConn),
	}
}
