package logger

import (
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"runtime"
	"time"
)

func SetupLogging() {
	detectTerminalAttached()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func SetupLoggingLevel(cfgLevel string) {
	if logLevel, err := zerolog.ParseLevel(cfgLevel); err == nil {
		zerolog.SetGlobalLevel(logLevel)
		return
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

//func detectTerminalAttached() {
//	if isatty.IsTerminal(os.Stdout.Fd()) && runtime.GOOS != "windows" {
//		log.Logger = log.Output(zerolog.ConsoleWriter{
//			Out:        os.Stdout,
//			TimeFormat: time.RFC3339,
//		})
//	}
//}

func detectTerminalAttached() {
	if isatty.IsTerminal(os.Stdout.Fd()) && runtime.GOOS != "windows" {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()
	}
}
