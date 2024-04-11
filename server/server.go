package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/realjoni17/Cdlsiet/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	router.Use(cors.New(config))

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run("0.0.0.0:" + s.port))
}
