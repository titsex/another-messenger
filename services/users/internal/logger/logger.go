package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var Print zerolog.Logger

func Init(level zerolog.Level) {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	Print = zerolog.New(consoleWriter).With().Timestamp().Caller().Logger().Level(level)
}
