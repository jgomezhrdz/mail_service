package server

import (
	"fmt"
	"log"
	mailing "mail_service/internal"
	"mail_service/internal/platform/server/handler/clientes"
	"mail_service/internal/platform/server/handler/health"
	"mail_service/internal/platform/server/middleware/cors"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	clienteReposiroty mailing.ClienteRepository
}

func New(host string, port uint, clienteReposiroty mailing.ClienteRepository) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		clienteReposiroty: clienteReposiroty,
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
	s.engine.GET("/courses", clientes.GetHandler(s.clienteReposiroty))
	s.engine.POST("/courses", clientes.CreateHandler(s.clienteReposiroty))
}
