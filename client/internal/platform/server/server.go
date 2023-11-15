package server

import (
	"fmt"
	"log"
	"mail_service/internal/platform/server/handler/clientes"
	"mail_service/internal/platform/server/handler/health"
	"mail_service/internal/platform/server/middleware/cors"
	cliente_services "mail_service/internal/services/cliente"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	clienteService cliente_services.ClienteService
}

func New(host string, port uint, clienteService cliente_services.ClienteService) Server {
	srv := Server{
		engine:         gin.New(),
		httpAddr:       fmt.Sprintf("%s:%d", host, port),
		clienteService: clienteService,
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
	s.engine.GET("/clientes", clientes.GetHandler(s.clienteService))
	s.engine.POST("/clientes", clientes.CreateHandler(s.clienteService))
}
