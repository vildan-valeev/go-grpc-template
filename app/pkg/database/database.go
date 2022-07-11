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
	"github.com/jackc/tern/migrate"
	"github.com/rs/zerolog"
)

var (
	ErrDSNRequired     = errors.New("dsn required")
	ErrFailedAppendPEM = errors.New("failed to append PEM")
)

type Config struct {
	DSN                   string `env:"DSN,required" envDefault:"postgres://postgres:postgres@localhost:5432/postgres"`
	MigrationsDir         string `env:"MIGRATION_MIGRATIONS_DIR,required" envDefault:"migrations"`
	MigrationConfig       string `env:"MIGRATION_CONFIG" envDefault:"migrations/tern.conf"`
	MigrationVersion      string `env:"MIGRATION_VERSION" envDefault:"last"`
	MigrationVersionTable string `env:"MIGRATION_VERSION_TABLE" envDefault:"public.schema_version"`
	MigrationAuto         bool   `env:"MIGRATION_AUTO" envDefault:"true"`
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

	// Start migrate.
	if db.config.MigrationAuto {
		conn, err := db.pool.Acquire(context.Background())
		if err != nil {
			db.logger.Fatal().Err(err).Msgf("Unable to acquire a database connection: %v", err)
		}
		db.migrate(conn.Conn())
		conn.Release()
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

func (db *DB) migrate(conn *pgx.Conn) {
	ctx := context.Background()
	migrator, err := migrate.NewMigrator(ctx, conn, db.config.MigrationVersionTable)
	if err != nil {
		db.logger.Fatal().Err(err).Msgf("Unable to create a migrator: %v", err)
	}

	err = migrator.LoadMigrations(db.config.MigrationsDir)
	if err != nil {
		db.logger.Fatal().Err(err).Msgf("Unable to load migrations: %v", err)
	}

	err = migrator.Migrate(ctx)
	if err != nil {
		db.logger.Fatal().Err(err).Msgf("Unable to migrate: %v", err)
	}

	ver, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		db.logger.Fatal().Err(err).Msgf("Unable to get current schema version: %v", err)
	}
	db.logger.Info().Msgf("Migration done. Current schema version: %v", ver)

}
