package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var logger *zap.SugaredLogger
var lock = &sync.Mutex{}

func initLogger() *zap.SugaredLogger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")
	return logger.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		lock.Lock()
		defer lock.Unlock()
		if logger == nil {
			// Creating single instance now.
			logger = initLogger()
		}
	}

	return logger
}
