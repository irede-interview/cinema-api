package threaterhdl

import (
	"github.com/gin-gonic/gin"
	"github.com/irede-interview/cinema-api/internal/adapters"
	threaterservice "github.com/irede-interview/cinema-api/internal/core/use-cases/threater"
)

type ThreaterHandler struct {
	service threaterservice.ThreaterService
}

func (cc *ThreaterHandler) SetUpRoutes(r *gin.Engine) {
	threaterRoutes := r.Group("/threaters")

	threaterRoutes.POST("/", cc.CreateThreaterAction)
	threaterRoutes.GET("/:threaterToken", cc.GetThreaterAction)
	threaterRoutes.GET("/", cc.ListThreatersAction)
	threaterRoutes.PUT("/:threaterToken", cc.UpdateThreaterAction)
	threaterRoutes.POST("/:threaterToken/inactivate", cc.UpdateThreaterAction)
}

func NewHandler(apt *adapters.Adapters) *ThreaterHandler {
	return &ThreaterHandler{
		service: *threaterservice.New(apt.Repositories, apt.Logger),
	}
}
