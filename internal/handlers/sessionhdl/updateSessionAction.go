package sessionhdl

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	sessionservice "github.com/irede-interview/cinema-api/internal/core/use-cases/session"
	"github.com/irede-interview/cinema-api/internal/utils"
)

type updateSessionDto struct {
	ThreaterToken   string    `json:"threater_token" binding:"required"`
	MovieToken      string    `json:"movie_token" binding:"required"`
	SessionDatetime time.Time `json:"session_datetime" binding:"required"`
}

func (cc *SessionHandler) UpdateSessionAction(context *gin.Context) {
	sessionToken := context.Param("sessionToken")
	var requestBody updateSessionDto
	if err := utils.BindAndValidate(context, &requestBody); err != nil {
		return
	}

	params := sessionservice.UpdateSessionParams{
		SessionToken:    sessionToken,
		ThreaterToken:   requestBody.ThreaterToken,
		MovieToken:      requestBody.MovieToken,
		SessionDatetime: requestBody.SessionDatetime,
	}

	err := cc.service.Update(params)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"success": true})
}
