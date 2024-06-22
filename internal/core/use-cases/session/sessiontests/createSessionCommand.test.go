package sessiontests

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/irede-interview/cinema-api/internal/core/domain"
	sessionservice "github.com/irede-interview/cinema-api/internal/core/use-cases/threater"
	"github.com/irede-interview/cinema-api/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListThreaterCommand_Execute(t *testing.T) {
	mockRepo := new(mocks.MockThreaterRepository)
	mockLogger := new(mocks.MockLogger)
	command := sessionservice.NewListThreatersCommand(mockRepo, mockLogger)

	expectedThreaters := []domain.Threater{
		{Token: uuid.MustParse("1"), Number: 1, Description: "Large Threater"},
		{Token: uuid.MustParse("2"), Number: 2, Description: "Small Threater"},
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("List").Return(expectedThreaters, nil)
		mockLogger.On("Info", mock.Anything).Twice()

		threaters, err := command.Execute()

		assert.NoError(t, err)
		assert.Equal(t, expectedThreaters, threaters)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("List").Return(nil, errors.New("database error"))
		mockLogger.On("Info", mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		threaters, err := command.Execute()

		assert.Error(t, err)
		assert.Nil(t, threaters)
		assert.EqualError(t, err, "error creating Threater: database error")
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
