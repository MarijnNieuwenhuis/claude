package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a new zap.SugaredLogger with the given log level.
//
// The log level should be one of the following: debug, info, warn, error, fatal, panic or dpanic.
// If an unknown log level is given, the log level will default to info.
func NewLogger(level string) *zap.SugaredLogger {
	c := zap.NewProductionConfig()
	c.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	c.DisableCaller = true
	c.EncoderConfig.MessageKey = "message"
	c.EncoderConfig.LevelKey = "level_name"
	c.EncoderConfig.TimeKey = "datetime"
	c.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	unknownLevel := false
	if l, err := zapcore.ParseLevel(level); err == nil {
		c.Level.SetLevel(l)
	} else {
		c.Level.SetLevel(zap.InfoLevel)
		unknownLevel = true
	}

	l, _ := c.Build()
	defer l.Sync()

	logger := l.Sugar()

	if unknownLevel {
		logger.Warnf("Could not set unknown log level '%s'. Defaulting to 'info'", level)
		logger.Info("Valid log levels are: debug, info, warn, error, fatal, panic and dpanic")
	}

	return logger
}
