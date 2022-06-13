package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-bolvanka/internal/config"
	"go-bolvanka/internal/domain/service"
	"go-bolvanka/internal/repository"
	"go-bolvanka/internal/transport/server"
	"go-bolvanka/pkg/database"
	"go-bolvanka/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.SetupLogging()

	log.Info().Msgf("Start App")

	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		sig := <-sigs
		log.Info().Msgf("Shutting down server. Reason: %s...", sig.String())
		cancel()
	}()

	// Instantiate a new type to represent our application.
	m := NewMain()

	// Execute program.
	if err := m.Run(ctx); err != nil {
		log.Error().Err(err).Msg("Run server error")

		_ = m.Close()

		os.Exit(1)
	}

	// Wait for CTRL-C.
	<-ctx.Done()

	// Clean up program.
	if err := m.Close(); err != nil {
		log.Error().Err(err).Msg("Shutting down server error")
		os.Exit(1)
	}

	log.Info().Msg("Bye!")
}

// Main represents the program.
type Main struct {
	// Config parsed config data.
	Config *config.Config
	// DB used by postgres service implementations.
	DB *database.DB
	// HTTP server for handling communication.
	Srv *server.Server
}

// NewMain returns a new instance of Main.
func NewMain() *Main {
	log.Info().Msg("Init config")
	cfg := config.New()

	logger.SetupLoggingLevel(cfg.LogLevel)

	return &Main{
		Config: cfg,
		DB:     database.New(cfg.Database, log.Logger),
	}
}

// Run executes the program. The configuration should already be set up before
// calling this function.
func (m *Main) Run(ctx context.Context) (err error) {
	// Then open the database. This will instantiate the connection
	// and execute any pending migration files.

	if err := m.DB.Open(ctx); err != nil {
		return err
	}

	repositories := repository.NewRepositories(m.DB)
	services := service.NewServices(service.Deps{
		Repos: repositories,
	})

	m.Srv = server.New(*m.Config, services)

	// Start the server.
	return m.Srv.Open()
}

// Close gracefully stops the program.
func (m *Main) Close() (err error) { //nolint
	if m.Srv != nil {
		_ = m.Srv.Close()
	}

	if m.DB != nil {
		_ = m.DB.Close()
	}

	return nil
}
