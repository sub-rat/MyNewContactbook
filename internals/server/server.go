package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MyNewContactbook/internals/controllers"
	"github.com/sub-rat/MyNewContactbook/internals/models"
	"gorm.io/gorm"
)

type server struct {
	C  *gin.Engine
	DB *gorm.DB
}

func GetServer() *server {
	return &server{
		C:  gin.Default(),
		DB: models.ConnectDatabase(),
	}
}

func (s *server) Run() {
	s.initRoutes()
	s.C.Run()
}

func (s *server) initRoutes() {
	// routes or Endpoints
	r := s.C
	r.GET("/", controllers.Welcome)
	r.GET("/ping", controllers.Ping)

	// contact endpoints
	r.GET("/contacts", controllers.GetAllContacts)
	r.POST("/contacts", controllers.CreateContact)

	r.PUT("/contacts/:id", controllers.UpdateContactById)
	r.GET("/contacts/:id", controllers.GetContactById)
	r.DELETE("/contacts/:id", controllers.DeleteContactsById)
}
