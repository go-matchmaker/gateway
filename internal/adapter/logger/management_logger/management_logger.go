package management_logger

import (
	"gateway/internal/adapter/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CreateLogger(selectedLevel int) *zap.Logger {
	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "storage/management/app.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     1,
		Compress:   true,
	})

	selected := logger.InitLogger(selectedLevel)
	level := zap.NewAtomicLevelAt(selected)

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, file, level),
	)

	return zap.New(core)
}
