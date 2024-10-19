package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateRequiredFields(ctx *gin.Context, body interface{}, fields map[string]string) bool {
	for fieldName, fieldValue := range fields {
		if fieldValue == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": fieldName + " is required"})
			return false
		}
	}
	return true
}
