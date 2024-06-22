package sessionhdl

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	sessionservice "github.com/irede-interview/cinema-api/internal/core/use-cases/session"
	"github.com/irede-interview/cinema-api/internal/utils"
)

type createSessionDto struct {
	ThreaterToken   string    `json:"threater_token" binding:"required"`
	MovieToken      string    `json:"movie_token" binding:"required"`
	SessionDatetime time.Time `json:"session_datetime" binding:"required"`
}

func (mh *SessionHandler) CreateSessionAction(context *gin.Context) {
	var requestBody createSessionDto
	if err := utils.BindAndValidate(context, &requestBody); err != nil {
		return
	}

	params := sessionservice.CreateSessionParams{
		ThreaterToken:   requestBody.ThreaterToken,
		MovieToken:      requestBody.MovieToken,
		SessionDatetime: requestBody.SessionDatetime,
	}

	createdSession, err := mh.service.Create(params)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": createdSession})
}
