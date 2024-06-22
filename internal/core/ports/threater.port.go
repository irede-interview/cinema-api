package ports

import "github.com/irede-interview/cinema-api/internal/core/domain"

type ThreaterRepository interface {
	Create(threater *domain.Threater) (*domain.Threater, error)
	Get(threaterToken string) (*domain.Threater, error)
	List() ([]domain.Threater, error)
	Update(threaterToUpdate *domain.Threater) error
	Inactivate(threaterToken string) error
}

type ThreaterService interface{}
