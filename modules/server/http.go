package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/raufhm/golang-project-convention/modules/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Router interface {
	Register(e *echo.Echo)
}

func RegisterRoutes(e *echo.Echo, routers []Router) {
	for _, router := range routers {
		router.Register(e)
	}
}

type Server struct {
	echo   *echo.Echo
	config *config.Config
	server *http.Server
}

func NewServer(cfg *config.Config) *Server {
	e := echo.New()
	// Add common middleware here

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	return &Server{
		echo:   e,
		config: cfg,
		server: httpServer,
	}
}

// GetEcho returns the echo instance for registering routes
func (s *Server) GetEcho() *echo.Echo {
	return s.echo
}

// Start starts the HTTP server with graceful shutdown
func (s *Server) Start() error {
	// Error channel to catch server startup errors
	errChan := make(chan error, 1)

	// Start server in goroutine
	go func() {
		s.echo.Logger.Info("starting server on ", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				errChan <- fmt.Errorf("server error: %v", err)
				return
			}
		}
	}()

	// Wait for interrupt signal or server error
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Block until signal or error
	select {
	case err := <-errChan:
		return fmt.Errorf("server failed to start: %v", err)
	case sig := <-quit:
		s.echo.Logger.Info("received shutdown signal: ", sig)
	}

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.echo.Logger.Info("shutting down server")
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %v", err)
	}

	s.echo.Logger.Info("server shutdown complete")
	return nil
}
