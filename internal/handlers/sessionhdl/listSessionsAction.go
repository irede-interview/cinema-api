package sessionhdl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sessionservice "github.com/irede-interview/cinema-api/internal/core/use-cases/session"
	"github.com/irede-interview/cinema-api/internal/utils"
)

func (mh *SessionHandler) ListSessionsAction(context *gin.Context) {
	var requestBody getSessionDto
	if err := utils.BindAndValidate(context, &requestBody); err != nil {
		return
	}

	pageStr := context.DefaultQuery("page", "1")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid page number", "success": false})
		return
	}

	params := sessionservice.ListSessionsParams{
		Page: page,
	}

	sessions, err := mh.service.List(params)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": sessions})
}
