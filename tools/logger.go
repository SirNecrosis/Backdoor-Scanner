package tools

import "go.uber.org/zap"

// Generic Global Logging Function
func Log() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return logger
}
