package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"leaderboard/adapters/rest"
	"leaderboard/internal/config"
	"leaderboard/internal/manager"
	"leaderboard/internal/storage"
	"leaderboard/internal/symbols"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// App contains the "global" components that are passed around
type App struct {
	logger  *zerolog.Logger
	config  *config.Manager
	storage *storage.Storage
	symbols *symbols.Manager
	events  *EventManger
	manager *manager.Manager
	// Channel for passing reload signals.
	signal chan os.Signal
	// Root context that is used to manage the application lifecycle
	// and is passed to all the components.
	ctx context.Context
}

var (
	logger = initLogger()
	conf   *config.Manager
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conf = config.New(logger)
	app := App{
		logger:  logger,
		ctx:     context.Background(),
		config:  conf,
		storage: storage.New(logger, conf),
		symbols: symbols.New(logger, conf),
	}
	app.logger.Info().Str("Version", app.config.GetServiceConfig().Version).Msg("Starting application")
	app.manager = manager.New(app.ctx, app.logger, app.symbols, app.storage)
	app.events = newEventManager(&app)
	app.events.start()
	server := rest.StartServer(app.logger, app.config, app.manager, app.symbols)

	// Wait for the interrupt or sigterm signal to gracefully shut down resources
	// within N seconds, or do a force shutdown.
	app.signal = make(chan os.Signal, 2)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be caught, so don't need to add it
	signal.Notify(app.signal, os.Interrupt, syscall.SIGTERM)
	for {
		sig := <-app.signal
		app.logger.Info().Str("signal", sig.String()).Msg("Received Signal to Shutdown")
		app.events.shutdown()
		app.manager.Shutdown()
		server.Shutdown(app.ctx)
		return
	}
}
