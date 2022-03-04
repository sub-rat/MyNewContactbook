package server

import (
	"github.com/gin-gonic/gin"
	contact2 "github.com/sub-rat/MyNewContactbook/internals/features/contact"
	user2 "github.com/sub-rat/MyNewContactbook/internals/features/user"
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
	contact2.RegisterRoutes(r, contact2.NewService(contact2.NewRepository(*s.DB)))

	user2.RegisterRoutes(r, user2.NewService(user2.NewRepository(*s.DB)))

}
