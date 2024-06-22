package threaterhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	threaterservice "github.com/irede-interview/cinema-api/internal/core/use-cases/threater"
)

type getThreaterDto struct {
	ThreaterToken string `json:"threater_token" binding:"required"`
}

func (mh *ThreaterHandler) GetThreaterAction(context *gin.Context) {
	threaterToken := context.Param("threaterToken")

	threater, err := mh.service.Get(threaterservice.GetThreaterParams{
		ThreaterToken: threaterToken,
	})
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": threater})
}
