package zerohook

import (
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

type SentryHook struct{}

func (h SentryHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level >= zerolog.WarnLevel {
		sentry.CaptureMessage(msg)
	}
}
