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

func StartServer(log *zerolog.Logger, config *config.Manager, manager *manager.Manager, symbols *symbols.Manager) *Server {
	r := gin.New()
	r.Use(middleware.StructuredLogger(log), gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowCredentials: true,
	}))
	handlers.SetupHandlers(r, log, config, symbols, manager)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Stack().Msg("Failed to start the server")
		}
	}()
	return &Server{server: srv, logger: log}
}

func (s *Server) Shutdown(ctx context.Context) {
	s.logger.Info().Msg("Shutting down the server")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error().Err(err).Stack().Msg("Failed to gracefully shutdown the server")
	}
	select {
	case <-ctx.Done():
		s.logger.Warn().Err(ctx.Err()).Msg("Timeout of 5 seconds while waiting for server to shutdown")
	}
	s.logger.Info().Msg("Server exiting")
}
