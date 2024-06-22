package sessionhdl

import (
	"github.com/gin-gonic/gin"
	"github.com/irede-interview/cinema-api/internal/adapters"
	sessionservice "github.com/irede-interview/cinema-api/internal/core/use-cases/session"
)

type SessionHandler struct {
	service sessionservice.SessionService
}

func (cc *SessionHandler) SetUpRoutes(r *gin.Engine) {
	sessionRoutes := r.Group("/sessions")

	sessionRoutes.POST("/", cc.CreateSessionAction)
	sessionRoutes.GET("/:sessionToken", cc.GetSessionAction)
	sessionRoutes.GET("/", cc.ListSessionsAction)
	sessionRoutes.PUT("/:sessionToken", cc.UpdateSessionAction)
	sessionRoutes.POST("/:sessionToken/inactivate", cc.UpdateSessionAction)
}

func NewHandler(apt *adapters.Adapters) *SessionHandler {
	return &SessionHandler{
		service: *sessionservice.New(apt.Repositories, apt.Logger),
	}
}
