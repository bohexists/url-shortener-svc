package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

// Logger - глобальная переменная логгера
var Logger *zap.Logger

// InitLogger инициализирует логгер Zap
func InitLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Удобный формат времени
	var err error
	Logger, err = config.Build()
	if err != nil {
		log.Fatalf("Не удалось инициализировать Zap логгер: %v", err)
	}
	defer Logger.Sync() // Синхронизация вывода (важно для flush логов)
}
