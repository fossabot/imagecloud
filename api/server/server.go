package server

import (
	"github.com/gin-gonic/gin"
	"github.com/songjiayang/imagecloud/api/handler"
	"github.com/songjiayang/imagecloud/internal/pkg/config"
)

type Server struct {
	*gin.Engine
	*handler.Handler
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Engine:  gin.Default(),
		Handler: handler.NewHandler(cfg),
	}
}

func (s *Server) Listen() error {
	// register routes
	s.GET("/:key", s.Image.Get)
	s.POST("/", s.Image.Post)
	s.GET("/", s.Pong)

	return s.Run()
}