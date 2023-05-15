package ports

import (
	"WB_L0/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	engine *gin.Engine
	port   string
}

func NewServer(port string, s service.Service) *Server {

	gin.SetMode(gin.ReleaseMode)

	server := &Server{
		engine: gin.Default(),
		port:   port,
	}
	Router(server.engine, s)

	return server
}
func (s *Server) Listen() {
	s.engine.Run(s.port)
}
func (s *Server) GetHandler() http.Handler {
	return s.engine
}
