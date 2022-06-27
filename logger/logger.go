package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log *zerolog.Logger

// InitLogger initialize the global Log logger on this package
// The purpose is to use this class logger as a common global logger
func InitLogger(level string) {
	var lev zerolog.Level

	switch level {
	case "info":
		lev = zerolog.InfoLevel
	case "debug":
		lev = zerolog.DebugLevel
	default:
		lev = zerolog.InfoLevel
	}

	logger := zerolog.New(os.Stderr).With().Logger().Level(lev)
	Log = &logger
}
