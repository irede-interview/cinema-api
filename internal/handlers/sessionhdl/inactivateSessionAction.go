package sessionhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sessionservice "github.com/irede-interview/cinema-api/internal/core/use-cases/session"
)

func (cc *SessionHandler) InactivateSessionAction(context *gin.Context) {
	sessionToken := context.Param("sessionToken")

	params := sessionservice.InactivateSessionParams{
		SessionToken: sessionToken,
	}

	err := cc.service.Inactivate(params)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"success": true})
}
