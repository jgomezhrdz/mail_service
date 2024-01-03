package server

import (
	"mail_service/internal/platform/server/handler/clientes"
	"mail_service/internal/platform/server/handler/health"
	"mail_service/internal/platform/server/middleware/cors"
)

func (s *Server) registerRoutes() {
	s.engine.Use(cors.Middleware())

	s.engine.GET("/health", health.CheckHandler())

	//CLIENTES
	clientesGroups := s.engine.Group("/clientes")
	clientesGroups.GET("", clientes.GetHandler(s.clienteService))
	clientesGroups.POST("", clientes.CreateHandler(s.clienteService))
	clientesGroups.PUT("", clientes.UpdateHandler(s.clienteService))
	clientesGroups.DELETE("", clientes.DeleteHandler(s.clienteService))
}
