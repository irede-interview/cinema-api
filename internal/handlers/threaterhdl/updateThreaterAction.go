package threaterhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	threaterservice "github.com/irede-interview/cinema-api/internal/core/use-cases/threater"
	"github.com/irede-interview/cinema-api/internal/utils"
)

type updateThreaterDto struct {
	Number      string `json:"number" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (cc *ThreaterHandler) UpdateThreaterAction(context *gin.Context) {
	threaterToken := context.Param("threaterToken")
	var requestBody updateThreaterDto
	if err := utils.BindAndValidate(context, &requestBody); err != nil {
		return
	}

	params := threaterservice.UpdateThreaterParams{
		ThreaterToken: threaterToken,
		Description:   requestBody.Description,
	}

	err := cc.service.Update(params)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"success": true})
}
