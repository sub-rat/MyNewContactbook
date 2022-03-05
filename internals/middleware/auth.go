package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckToken(context *gin.Context) {
	authKey := context.GetHeader("Authorization")
	if authKey != "ABD!23@@#" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token Required"})
		return
	}
	context.Next()
}
