package sessionservice

import (
	"fmt"

	"github.com/irede-interview/cinema-api/internal/core/domain"
	"github.com/irede-interview/cinema-api/internal/core/ports"
	"github.com/irede-interview/cinema-api/pkg/logger"
)

type ListSessionsCommand struct {
	sessionRepository ports.SessionRepository
	logger            logger.Provider
}

func NewListSessionsCommand(SessionRepo ports.SessionRepository, logger logger.Provider) *ListSessionsCommand {
	return &ListSessionsCommand{
		sessionRepository: SessionRepo,
		logger:            logger,
	}
}

func (cmd *ListSessionsCommand) Execute() ([]domain.Session, error) {
	cmd.logger.Info("ListSessionsCommand initiated")

	sessions, err := cmd.sessionRepository.List()
	if err != nil {
		cmd.logger.Error("ListSessionsCommand initiated", err)
		return nil, fmt.Errorf("error creating Session: %w", err)
	}

	cmd.logger.Info("ListSessionsCommand finished")

	return sessions, nil
}
