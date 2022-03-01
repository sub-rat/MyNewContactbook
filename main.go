package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MyNewContactbook/controllers"
	"github.com/sub-rat/MyNewContactbook/models"
)

func main() {
	fmt.Println("Starting ContactBook Api")
	r := gin.Default()

	//Database connection
	models.ConnectDatabase()

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

	r.PUT("/contacts", controllers.UpdateContact)

	r.GET("/contacts/:id", controllers.GetContactById)
	r.DELETE("/contacts/:id", controllers.DeleteContactsById)
}
