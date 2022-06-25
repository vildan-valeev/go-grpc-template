package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"go-grpc-template/pkg/logger/pgxadapter"
	"go-grpc-template/pkg/logger/zerohook"

	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

var (
	ErrDSNRequired     = errors.New("dsn required")
	ErrFailedAppendPEM = errors.New("failed to append PEM")
)

type Config struct {
	DSN string `env:"DSN,required" envDefault:"postgres://postgres:postgres@localhost:5432/postgres"`
}

// DB represents the database connection.
type DB struct {
	config Config
	pool   *pgxpool.Pool
	logger zerolog.Logger
	level  pgx.LogLevel

	// Returns the current time. Defaults to time.Now().
	// Can be mocked for tests.
	Now func() time.Time
}

// New returns a new instance of DB associated with the given datasource name.
func New(cfg Config, log zerolog.Logger) *DB {
	db := &DB{
		pool:   nil,
		logger: log.With().Str("module", "pgx").Logger().Hook(zerohook.SentryHook{}),
		config: cfg,
		Now:    time.Now,
	}

	return db
}

// Open opens the database connection.
func (db *DB) Open(ctx context.Context) (err error) {
	// Ensure a DSN is set before attempting to open the database.
	log.Info().Msg("Init DB")
	if db.config.DSN == "" {
		return ErrDSNRequired
	}

	poolConfig, err := pgxpool.ParseConfig(db.config.DSN)
	if err != nil {
		return err
	}

	db.level, err = pgx.LogLevelFromString(db.logger.GetLevel().String())
	if err != nil {
		db.level = pgx.LogLevelWarn
	}

	poolConfig.ConnConfig.Logger = pgxadapter.NewLogger(db.logger)
	poolConfig.ConnConfig.LogLevel = db.level

	// Connect to the database.
	db.pool, err = pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err) //nolintlint:goerr113
	}

	return nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	if db.pool != nil {
		db.pool.Close()
	}

	return nil
}

func (db *DB) Pool() *pgxpool.Pool {
	return db.pool
}
