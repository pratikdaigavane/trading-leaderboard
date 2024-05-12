package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"leaderboard/adapters/rest"
	"leaderboard/internal/config"
	"leaderboard/internal/storage"
	"leaderboard/internal/symbols"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	ctx     context.Context
	signal  chan os.Signal
	logger  *zerolog.Logger
	config  *config.Manager
	storage *storage.Storage
	symbols *symbols.Manager
	events  *EventManger
}

func main() {
	time.Local = time.UTC
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := App{
		logger: initLogger(),
		ctx:    context.Background(),
	}
	app.config = config.New(app.logger)
	app.logger.Info().Str("Version", app.config.GetServiceConfig().Version).Msg("Starting application")
	app.storage = storage.New(app.logger, app.config)
	app.symbols = symbols.New(app.logger, app.config)
	app.events = newEventManager(&app)
	app.events.start()
	/*trade := models.Trade{
		Symbol:   "AAPL",
		Quantity: 100,
		TraderId: "dfdsfdsfsd",
		ID:       "sdfdfsdf",
	}
	err = app.storage.Operations.Add(&trade)
	if err != nil {
		app.logger.Error().Err(err).Stack().Msg("Failed to add trade")
		return
	}
	stats, _ := app.storage.Operations.GetSortedList("AdAPL", 10)
	app.logger.Info().Interface("stats", stats).Msg("Stats")*/
	server := rest.StartServer(app.logger, app.config, app.storage, app.symbols)
	app.signal = make(chan os.Signal, 2)
	signal.Notify(app.signal, os.Interrupt, syscall.SIGTERM)
	for {
		sig := <-app.signal
		app.logger.Info().Str("signal", sig.String()).Msg("Received Signal to Shutdown")
		app.events.shutdown()
		server.Shutdown(app.ctx)
		return
	}
}
