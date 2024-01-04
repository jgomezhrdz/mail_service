package server

import (
	_ "mail_service/swagger"

	"mail_service/internal/platform/server/handler/clientes"
	"mail_service/internal/platform/server/handler/health"
	"mail_service/internal/platform/server/middleware/cors"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) registerRoutes() {
	s.engine.Use(cors.Middleware())

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	s.engine.GET("/health", health.CheckHandler())

	//CLIENTES
	clientesGroups := s.engine.Group("/clientes")
	clientesGroups.GET("", clientes.GetHandler(s.clienteService))
	clientesGroups.POST("", clientes.CreateHandler(s.clienteService))
	clientesGroups.PUT("", clientes.UpdateHandler(s.clienteService))
	clientesGroups.DELETE("", clientes.DeleteHandler(s.clienteService))
}
