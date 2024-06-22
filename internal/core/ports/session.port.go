package ports

import (
	"github.com/irede-interview/cinema-api/internal/core/domain"
)

type SessionRepository interface {
	Create(session *domain.Session) (*domain.Session, error)
	Get(sessionToken string) (*domain.Session, error)
	List() ([]domain.Session, error)
	Update(sessionToUpdate *domain.Session) error
	Inactivate(movieToken string) error
}

type SessionService interface{}
