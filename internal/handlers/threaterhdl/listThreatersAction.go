package threaterhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irede-interview/cinema-api/internal/utils"
)

func (mh *ThreaterHandler) ListThreatersAction(context *gin.Context) {
	var requestBody getThreaterDto
	if err := utils.BindAndValidate(context, &requestBody); err != nil {
		return
	}

	threaters, err := mh.service.List()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": threaters})
}
