package sessionhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sessionservice "github.com/irede-interview/cinema-api/internal/core/use-cases/session"
)

type getSessionDto struct {
	SessionToken string `json:"session_token" binding:"required"`
}

func (mh *SessionHandler) GetSessionAction(context *gin.Context) {
	sessionToken := context.Param("sessionToken")

	session, err := mh.service.Get(sessionservice.GetSessionParams{
		SessionToken: sessionToken,
	})
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": session})
}
