package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MyNewContactbook/controller"
)

func main() {
	fmt.Println("Starting ContactBook Api")
	r := gin.Default()
	initRoutes(r)
	// Starting Server, shortcut of http.ListenAndServer
	r.Run()
}

func initRoutes(r *gin.Engine) {
	// routes or Endpoints
	r.GET("/", controller.Welcome)
	r.GET("/ping", controller.Ping)

	// contact endpoints
	r.GET("/contacts", controller.GetAllContacts)
	r.POST("/contacts")
	r.DELETE("/contacts")
	r.PUT("/contacts")
}
