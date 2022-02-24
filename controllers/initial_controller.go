package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome Handler function for endpoints
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to ContactBook API",
	})
}

// Ping Handler function for endpoints
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
