package movierepo

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
	"github.com/irede-interview/cinema-api/internal/core/domain"
)

type MovieRepository struct {
	db *dbr.Session
}

func New(db *dbr.Session) *MovieRepository {
	repo := &MovieRepository{
		db: db,
	}

	return repo
}

func (r *MovieRepository) Create(movie *domain.Movie) (*domain.Movie, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.RollbackUnlessCommitted()

	movie.Token = uuid.New()
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	_, err = tx.InsertInto("movies").Columns(
		"token",
		"name",
		"director",
		"duration",
		"created_at",
		"updated_at",
	).Record(movie).Exec()
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return movie, nil
}

func (r *MovieRepository) List() ([]domain.Movie, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	var movies []domain.Movie
	_, err = tx.
		Select("*").
		From("movies").
		Load(&movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MovieRepository) Get(MovieToken string) (*domain.Movie, error) {
	var movie domain.Movie

	result, err := r.
		db.
		Select("*").
		From("movies").
		Where("token = ?", MovieToken).
		Load(&movie)

	if err != nil {
		return nil, err
	}

	if result == 0 {
		return nil, fmt.Errorf("Movie not found")
	}

	return &movie, nil
}

func (r *MovieRepository) Update(movieToUpdate *domain.Movie) error {
	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	defer tx.RollbackUnlessCommitted()

	_, err = tx.
		Update("movies").
		SetMap(map[string]interface{}{
			"name":     movieToUpdate.Name,
			"director": movieToUpdate.Director,
			"duration": movieToUpdate.Duration,
		}).
		Where(dbr.Eq("token", movieToUpdate.Token)).
		Exec()

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *MovieRepository) Inactivate(movieToken string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.RollbackUnlessCommitted()

	_, err = tx.
		Update("movies").
		Set("active", false).
		Where(dbr.Eq("token", movieToken)).
		Exec()

	if err != nil {
		return err
	}

	return tx.Commit()
}
