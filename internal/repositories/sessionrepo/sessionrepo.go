package sessionrepo

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
	"github.com/irede-interview/cinema-api/internal/core/domain"
)

type SessionRepository struct {
	db *dbr.Session
}

func New(db *dbr.Session) *SessionRepository {
	repo := &SessionRepository{
		db: db,
	}

	return repo
}

func (r *SessionRepository) Create(session *domain.Session) (*domain.Session, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.RollbackUnlessCommitted()

	session.Token = uuid.New()
	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()

	_, err = tx.InsertInto("sessions").Columns(
		"token",
		"movie_token",
		"thread_token",
		"session_datetime",
		"created_at",
		"updated_at",
	).Record(session).Exec()
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return session, nil
}

func (r *SessionRepository) List(page int) ([]domain.Session, error) {
	if page < 1 {
		page = 1
	}
	pageSize := 6
	offset := (page - 1) * pageSize

	var sessions []domain.Session
	tx, err := r.db.NewSession(nil).Begin()
	if err != nil {
		return nil, err
	}
	defer tx.RollbackUnlessCommitted()

	_, err = tx.Select("*").
		From("sessions").
		OrderBy("session_datetime DESC").
		Limit(uint64(pageSize)).
		Offset(uint64(offset)).
		Load(&sessions)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return sessions, nil
}

func (r *SessionRepository) Get(sessionToken string) (*domain.Session, error) {
	var session domain.Session

	result, err := r.
		db.
		Select("*").
		From("sessions").
		Where("token = ?", sessionToken).
		Load(&session)

	if err != nil {
		return nil, err
	}

	if result == 0 {
		return nil, fmt.Errorf("session not found")
	}

	return &session, nil
}

func (r *SessionRepository) Update(sessionToUpdate *domain.Session) error {
	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	defer tx.RollbackUnlessCommitted()

	_, err = tx.
		Update("sessions").
		SetMap(map[string]interface{}{
			"session_datetime": sessionToUpdate.SessionDatetime,
		}).
		Where(dbr.Eq("token", sessionToUpdate.Token)).
		Exec()

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *SessionRepository) Inactivate(sessionToken string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.RollbackUnlessCommitted()

	_, err = tx.
		Update("sessions").
		Set("active", false).
		Where(dbr.Eq("token", sessionToken)).
		Exec()

	if err != nil {
		return err
	}

	return tx.Commit()
}
