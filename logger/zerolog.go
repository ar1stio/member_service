package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func NewLogger() *zerolog.Logger {

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano}
	log := zerolog.New(output).With().Timestamp().Logger()

	return &log
}

func Error() *zerolog.Event {

	return NewLogger().Error().Str("app", "member-service")
}

func Info() *zerolog.Event {

	return NewLogger().Info().Str("app", "member-service")
}

func Warn() *zerolog.Event {

	return NewLogger().Warn().Str("app", "member-service")
}

func Trace() *zerolog.Event {
	return NewLogger().Trace().Str("app", "member-service")
}
