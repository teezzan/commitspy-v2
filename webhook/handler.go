package webhook

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/response"
	"github.com/teezzan/commitspy-v2/validator"
)

type EventHandlers struct{}

func (ctrl EventHandlers) Github(c *gin.Context) {

	var json validator.URIProjectID

	if err := c.ShouldBindUri(&json); err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.WriteSuccess(c, http.StatusOK, gin.H{})
}
