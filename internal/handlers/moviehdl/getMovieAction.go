package moviehdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	movieservice "github.com/irede-interview/cinema-api/internal/core/use-cases/movie"
)

type getMovieDto struct {
	MovieToken string `json:"movie_token" binding:"required"`
}

func (mh *MovieHandler) GetMovieAction(context *gin.Context) {
	movieToken := context.Param("movieToken")

	movie, err := mh.service.Get(movieservice.GetMovieParams{
		MovieToken: movieToken,
	})
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": movie})
}
