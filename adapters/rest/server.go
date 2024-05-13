package rest

import (
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"leaderboard/adapters/rest/handlers"
	"leaderboard/adapters/rest/middleware"
	"leaderboard/internal/config"
	"leaderboard/internal/manager"
	"leaderboard/internal/symbols"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	logger *zerolog.Logger
}

// StartServer initialises new gin instance, registers middleware and handlers and starts the REST server in a separate goroutine.
func StartServer(log *zerolog.Logger, config *config.Manager, manager *manager.Manager, symbols *symbols.Manager) *Server {
	r := gin.New()
	r.Use(middleware.StructuredLogger(log), gin.Recovery())

	// CORS middleware to allow all origins.
	// This is just for demo purposes, in production, this should be restricted to only the required origins
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowCredentials: true,
	}))
	handlers.SetupHandlers(r, log, config, symbols, manager)
	srv := &http.Server{
		Addr:    config.GetServiceConfig().HttpServerAddr,
		Handler: r.Handler(),
	}
	// Start the server in a separate goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Stack().Msg("Failed to start the server")
		}
	}()
	return &Server{server: srv, logger: log}
}

// Shutdown gracefully shuts down the REST server
func (s *Server) Shutdown(ctx context.Context) {
	s.logger.Info().Msg("Shutting down the server")
	// A channel that will shut down the server with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error().Err(err).Stack().Msg("Failed to gracefully shutdown the server")
	}
	// If the context is done (time out exceeded), force shutdown the server
	select {
	case <-ctx.Done():
		s.logger.Warn().Err(ctx.Err()).Msg("Timeout of 5 seconds while waiting for server to shutdown")
	}
	s.logger.Info().Msg("Server exiting")
}
