package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MyNewContactbook/pkg/utils"
	"net/http"
)

func CheckToken(context *gin.Context) {
	token, err := utils.VerifyJwtToken(context)
	if err != nil || !token.Valid {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println("userId", claims["id"])
	context.Set("id", claims["id"])
	context.Next()
}
