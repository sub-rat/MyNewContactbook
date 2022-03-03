package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MyNewContactbook/internals/contact"
	dbconnect "github.com/sub-rat/MyNewContactbook/pkg"
	"gorm.io/gorm"
)

type server struct {
	C  *gin.Engine
	DB *gorm.DB
}

func GetServer() *server {
	return &server{
		C:  gin.Default(),
		DB: dbconnect.ConnectDatabase(),
	}
}

func (s *server) Run() {
	s.initRoutes()
	s.C.Run()
}

func (s *server) initRoutes() {
	// routes or Endpoints
	r := s.C

	// Contact Routes
	contact.RegisterRoutes(r, contact.NewService(contact.NewRepository(*s.DB)))
}
