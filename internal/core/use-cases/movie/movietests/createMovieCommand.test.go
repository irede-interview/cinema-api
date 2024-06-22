package movieservice

import (
	"errors"
	"testing"

	"github.com/irede-interview/cinema-api/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMovieRepository struct {
	mock.Mock
}

func (m *MockMovieRepository) Create(movie *domain.Movie) (*domain.Movie, error) {
	args := m.Called(movie)
	return args.Get(0).(*domain.Movie), args.Error(1)
}

type MockLogger struct {
	mock.Mock
}

func (l *MockLogger) Info(msg string, params ...interface{}) {
	l.Called(msg, params)
}

func (l *MockLogger) Error(msg string, params ...interface{}) {
	l.Called(msg, params)
}

func TestCreateMovieCommand_Execute(t *testing.T) {
	mockRepo := new(MockMovieRepository)
	mockLogger := new(MockLogger)
	command := NewCreateMovieCommand(mockRepo, mockLogger)

	params := CreateMovieParams{
		Name:     "Inception",
		Director: "Christopher Nolan",
		Duration: 148,
	}

	expectedMovie := &domain.Movie{
		Name:     params.Name,
		Director: params.Director,
		Duration: params.Duration,
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*domain.Movie")).Return(expectedMovie, nil)
		mockLogger.On("Info", mock.Anything, mock.Anything).Twice()

		result, err := command.Execute(params)

		assert.NoError(t, err)
		assert.Equal(t, expectedMovie, result)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*domain.Movie")).Return(nil, errors.New("failed to create movie"))
		mockLogger.On("Info", mock.Anything, mock.Anything).Once()
		mockLogger.On("Error", mock.Anything, mock.Anything).Once()

		result, err := command.Execute(params)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})
}
