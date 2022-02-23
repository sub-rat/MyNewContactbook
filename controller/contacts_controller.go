package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllContacts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getting all contacts",
	})
}
