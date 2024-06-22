package moviehdl

import (
	"github.com/gin-gonic/gin"
	"github.com/irede-interview/cinema-api/internal/adapters"
	movieservice "github.com/irede-interview/cinema-api/internal/core/use-cases/movie"
	"github.com/irede-interview/cinema-api/pkg/logger"
)

type MovieHandler struct {
	service movieservice.MovieService
	logger  logger.Provider
}

func (cc *MovieHandler) SetUpRoutes(r *gin.Engine) {
	movieRoutes := r.Group("/movies")

	movieRoutes.POST("/", cc.CreateMovieAction)
	movieRoutes.GET("/:movieToken", cc.GetMovieAction)
	movieRoutes.GET("/", cc.ListMoviesAction)
	movieRoutes.PUT("/:movieToken", cc.UpdateMovieAction)
	movieRoutes.POST("/:movieToken/inactivate", cc.UpdateMovieAction)
}

func NewHandler(apt *adapters.Adapters) *MovieHandler {
	return &MovieHandler{
		service: *movieservice.New(apt.Repositories, apt.Logger),
	}
}
