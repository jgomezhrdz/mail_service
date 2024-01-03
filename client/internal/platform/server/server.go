package server

import (
	"context"
	"fmt"
	"log"
	cron_scheduler "mail_service/internal/platform/cron"
	cliente_services "mail_service/internal/services/cliente_services"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	cronScheduler *cron.Cron

	shutdownTimeout time.Duration

	clienteService cliente_services.ClienteService
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, clienteService cliente_services.ClienteService) (context.Context, Server) {
	srv := Server{
		engine:          gin.Default(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		shutdownTimeout: shutdownTimeout,
		clienteService:  clienteService,
		cronScheduler:   cron_scheduler.InitializeCron(),
	}

	srv.registerRoutes()
	srv.registerJobs()

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
