// Package pgxadapter provides a logger that writes to a github.com/rs/zerolog.
package pgxadapter

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

// NewLogger accepts a zerolog.Logger as input and returns a new custom pgx
// logging fascade as output.
func NewLogger(logger zerolog.Logger) *Logger {
	return &Logger{
		logger: logger.With().Logger(),
	}
}

func (pl *Logger) Log(_ context.Context, lvl pgx.LogLevel, msg string, data map[string]interface{}) {
	var zlevel zerolog.Level

	switch lvl {
	case pgx.LogLevelNone:
		zlevel = zerolog.NoLevel
	case pgx.LogLevelError:
		zlevel = zerolog.ErrorLevel
	case pgx.LogLevelWarn:
		zlevel = zerolog.WarnLevel
	case pgx.LogLevelInfo:
		zlevel = zerolog.DebugLevel
	case pgx.LogLevelDebug:
		zlevel = zerolog.DebugLevel
	default:
		zlevel = zerolog.DebugLevel
	}

	pgxlog := pl.logger.With().Fields(data).Logger()
	pgxlog.WithLevel(zlevel).Msg(msg)
}
