package server

import (
	"fmt"
	"log"
	"mail_service/internal/platform/server/handler/clientes"
	"mail_service/internal/platform/server/handler/health"
	"mail_service/internal/platform/server/middleware/cors"
	"mail_service/internal/platform/storage/mysql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	clienteReposiroty mysql.ClienteRepository
}

func New(host string, port uint, courseRepository mysql.ClienteRepository) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.Use(cors.Middleware())

	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", clientes.CreateHandler(s.clienteReposiroty))
}
