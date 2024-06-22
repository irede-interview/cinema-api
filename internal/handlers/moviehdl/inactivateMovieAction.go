package moviehdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	movieservice "github.com/irede-interview/cinema-api/internal/core/use-cases/movie"
)

func (cc *MovieHandler) InactivateMovieAction(context *gin.Context) {
	movieToken := context.Param("movieToken")

	params := movieservice.InactivateMovieParams{
		MovieToken: movieToken,
	}

	err := cc.service.Inactivate(params)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"success": true})
}
