package log

import (
	"context"
	stdLog "log"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Level is used to define the minimum logging level
// Check the leveling at https://github.com/rs/zerolog/blob/v1.20.0/log.go#L110
var Level zerolog.Level

// SetLevel receives the environment var with the level info, parses it and defines the value as global var
func SetLevel(levelString string) {
	level, _ := strconv.Atoi(levelString)
	Level = zerolog.Level(level)
}

// New defines a new logger injected in the context to be used in the entire stack
// already injecting a tracing ID.
func New(ctx context.Context, id string) context.Context {
	zerolog.SetGlobalLevel(Level)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano})
	logger := log.With().Str("request_id", id).Logger()
	ctx = logger.WithContext(ctx)
	return ctx
}

// Insert inserts a new field to be injected in the logger context to be used in the entire stack
func Insert(ctx context.Context, key, value string) context.Context {
	logger := log.Ctx(ctx).With().Str(key, value).Logger()
	ctx = logger.WithContext(ctx)
	return ctx
}

// Debug is a wrapper for logging at debug level with the request ID
func Debug(ctx context.Context, message string) {
	if ctx == nil {
		stdLog.Println(message)
		return
	}
	if e := log.Debug(); !e.Enabled() {
		return
	}
	log.Ctx(ctx).Debug().Msg(message)
}

// Debugf is a wrapper for formatted logging at debug level with the request ID
func Debugf(ctx context.Context, message string, v ...interface{}) {
	if ctx == nil {
		stdLog.Println(message)
		return
	}
	if e := log.Debug(); !e.Enabled() {
		return
	}
	log.Ctx(ctx).Debug().Msgf(message, v...)
}

// Info is a wrapper for logging at information level with the request ID
func Info(ctx context.Context, message string) {
	if ctx == nil {
		stdLog.Println(message)
		return
	}
	if e := log.Info(); !e.Enabled() {
		return
	}
	log.Ctx(ctx).Info().Msg(message)
}

// Infof is a wrapper for formatted logging at information level with the request ID
func Infof(ctx context.Context, message string, v ...interface{}) {
	if ctx == nil {
		stdLog.Println(message)
		return
	}
	if e := log.Info(); !e.Enabled() {
		return
	}
	log.Ctx(ctx).Info().Msgf(message, v...)
}

// Warning is a wrapper for logging at warning level with the request ID
func Warning(ctx context.Context, message string) {
	if ctx == nil {
		stdLog.Println(message)
		return
	}
	if e := log.Warn(); !e.Enabled() {
		return
	}
	log.Ctx(ctx).Warn().Msg(message)
}

// Warningf is a wrapper for formatted logging at warning level with the request ID
func Warningf(ctx context.Context, message string, v ...interface{}) {
	if ctx == nil {
		stdLog.Println(message)
		return
	}
	if e := log.Warn(); !e.Enabled() {
		return
	}
	log.Ctx(ctx).Warn().Msgf(message, v...)
}

// Error is a wrapper for logging for errors with the request ID
func Error(ctx context.Context, err error, message string) {
	if ctx == nil {
		stdLog.Println(message)
		return
	}
	if e := log.Err(nil); !e.Enabled() {
		return
	}
	log.Ctx(ctx).Err(err).Msg(message)
}

// Errorf is a wrapper for formatted logging for errors with the request ID
func Errorf(ctx context.Context, err error, message string, v ...interface{}) {
	if ctx == nil {
		stdLog.Println(message)
		return
	}
	if e := log.Err(nil); !e.Enabled() {
		return
	}
	log.Ctx(ctx).Err(err).Msgf(message, v...)
}

// Fatal is a wrapper for logging at fatal level with the request ID
func Fatal(ctx context.Context, message string) {
	if ctx == nil {
		stdLog.Println(message)
		return
	}
	if e := log.Fatal(); !e.Enabled() {
		return
	}
	log.Ctx(ctx).Fatal().Msg(message)
}
