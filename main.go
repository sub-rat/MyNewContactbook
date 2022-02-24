package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MyNewContactbook/controllers"
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
	r.GET("/", controllers.Welcome)
	r.GET("/ping", controllers.Ping)

	// contact endpoints
	r.GET("/contacts", controllers.GetAllContacts)
	r.POST("/contacts", controllers.CreateContact)
	r.DELETE("/contacts", controllers.DeleteContacts)
	r.PUT("/contacts", controllers.UpdateContact)

	r.GET("/contacts/:id", controllers.GetContactById)
}
