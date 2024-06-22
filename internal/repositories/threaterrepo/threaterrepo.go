package threaterrepo

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
	"github.com/irede-interview/cinema-api/internal/core/domain"
)

type ThreaterRepository struct {
	db *dbr.Session
}

func New(db *dbr.Session) *ThreaterRepository {
	repo := &ThreaterRepository{
		db: db,
	}

	return repo
}

func (t *ThreaterRepository) Create(threater *domain.Threater) (*domain.Threater, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.RollbackUnlessCommitted()

	threater.Token = uuid.New()
	threater.CreatedAt = time.Now()
	threater.UpdatedAt = time.Now()

	_, err = tx.InsertInto("threaters").Columns(
		"token",
		"movie_token",
		"thread_token",
		"threater_datetime",
		"created_at",
		"updated_at",
	).Record(threater).Exec()
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return threater, nil
}

func (t *ThreaterRepository) List() ([]domain.Threater, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	var threaters []domain.Threater
	_, err = tx.
		Select("*").
		From("threaters").
		Load(&threaters)
	if err != nil {
		return nil, err
	}

	return threaters, nil
}

func (t *ThreaterRepository) Get(threaterToken string) (*domain.Threater, error) {
	var threater domain.Threater

	result, err := t.
		db.
		Select("*").
		From("threaters").
		Where("token = ?", threaterToken).
		Load(&threater)

	if err != nil {
		return nil, err
	}

	if result == 0 {
		return nil, fmt.Errorf("threater not found")
	}

	return &threater, nil
}

func (t *ThreaterRepository) Update(threaterToUpdate *domain.Threater) error {
	tx, err := t.db.Begin()

	if err != nil {
		return err
	}

	defer tx.RollbackUnlessCommitted()

	_, err = tx.
		Update("threaters").
		SetMap(map[string]interface{}{
			"number":      threaterToUpdate.Number,
			"description": threaterToUpdate.Description,
		}).
		Where(dbr.Eq("token", threaterToUpdate.Token)).
		Exec()

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ThreaterRepository) Inactivate(threaterToken string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.RollbackUnlessCommitted()

	_, err = tx.
		Update("threaters").
		Set("active", false).
		Where(dbr.Eq("token", threaterToken)).
		Exec()

	if err != nil {
		return err
	}

	return tx.Commit()
}
