package server

import (
	"context"
	"fmt"
	"log"
	"mail_service/internal/platform/server/handler/clientes"
	"mail_service/internal/platform/server/handler/health"
	"mail_service/internal/platform/server/middleware/cors"
	cliente_services "mail_service/internal/services/cliente"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	clienteService cliente_services.ClienteService
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, clienteService cliente_services.ClienteService) (context.Context, Server) {
	srv := Server{
		engine:          gin.New(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		shutdownTimeout: shutdownTimeout,
		clienteService:  clienteService,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}

func (s *Server) registerRoutes() {
	s.engine.Use(cors.Middleware())

	s.engine.GET("/health", health.CheckHandler())
	s.engine.GET("/clientes", clientes.GetHandler(s.clienteService))
	s.engine.POST("/clientes", clientes.CreateHandler(s.clienteService))
}
