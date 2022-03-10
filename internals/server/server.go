package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MyNewContactbook/internals/features/contact"
	"github.com/sub-rat/MyNewContactbook/internals/features/user"
	"github.com/sub-rat/MyNewContactbook/pkg/db/postgres"
	"gorm.io/gorm"
	"log"
)

type server struct {
	C  *gin.Engine
	DB *gorm.DB
}

func GetServer() *server {
	return &server{
		C:  gin.Default(),
		DB: postgres.ConnectDatabase(),
	}
}

func (s *server) Run() {
	s.initRoutes()
	log.Fatal(s.C.Run())
}

func (s *server) initRoutes() {
	// routes or Endpoints
	r := s.C

	// Contact Routes
	contact.RegisterRoutes(r, contact.NewService(contact.NewRepository(*s.DB)))

	user.RegisterRoutes(r, user.NewService(user.NewRepository(*s.DB)))

}
