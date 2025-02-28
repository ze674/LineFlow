package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(production bool) (*zap.SugaredLogger, error) {
	config := zap.NewDevelopmentConfig()
	if production {
		config = zap.NewProductionConfig()
	}

	// Настройка формата времени
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
