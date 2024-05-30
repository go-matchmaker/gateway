package logger

import "go.uber.org/zap/zapcore"

//ip method level timestamp msg+statuscode

const (
	DebugLevel = zapcore.DebugLevel + 2
	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel + 2
	// WarnLevel logs are more important than Info, but don't need individual human review.
	WarnLevel = zapcore.WarnLevel + 2
	// ErrorLevel logs are high-priority. If an application is running smoothly, it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel + 2
	// DPanicLevel logs are particularly important errors. In development the logger panics after writing the message.
	DPanicLevel = zapcore.DPanicLevel + 2
	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel + 2
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel + 2
)

func InitLogger(level int) zapcore.Level {
	var loggerLevel zapcore.Level
	switch level {
	case int(DebugLevel):
		loggerLevel = zapcore.DebugLevel
	case int(InfoLevel):
		loggerLevel = zapcore.InfoLevel
	case int(WarnLevel):
		loggerLevel = zapcore.WarnLevel
	case int(ErrorLevel):
		loggerLevel = zapcore.ErrorLevel
	case int(DPanicLevel):
		loggerLevel = zapcore.DPanicLevel
	case int(PanicLevel):
		loggerLevel = zapcore.PanicLevel
	case int(FatalLevel):
		loggerLevel = zapcore.FatalLevel
	default:
		loggerLevel = zapcore.InfoLevel
	}
	return loggerLevel
}
