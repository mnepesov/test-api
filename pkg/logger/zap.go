package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(name, env, level string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	
	err := config.Level.UnmarshalText([]byte(level))
	if err != nil || len(level) == 0 {
		config.Level.SetLevel(zap.DebugLevel)
	}
	
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	
	logger = logger.With(
		zap.String("name", name),
		zap.String("env", env),
		zap.String("log_level", level),
	)
	
	return logger, nil
}
