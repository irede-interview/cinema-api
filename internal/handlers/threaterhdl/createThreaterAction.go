package threaterhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	threaterservice "github.com/irede-interview/cinema-api/internal/core/use-cases/threater"
	"github.com/irede-interview/cinema-api/internal/utils"
)

type createThreaterDto struct {
	Number      int    `json:"number" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (mh *ThreaterHandler) CreateThreaterAction(context *gin.Context) {
	var requestBody createThreaterDto
	if err := utils.BindAndValidate(context, &requestBody); err != nil {
		return
	}

	params := threaterservice.CreateThreaterParams{
		Number:      requestBody.Number,
		Description: requestBody.Description,
	}

	createdThreater, err := mh.service.Create(params)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "sumhess": false})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"sumhess": true, "data": createdThreater})
}
