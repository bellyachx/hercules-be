package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(msg string, fields ...any)
	Debugf(msg string, fields ...any)
	Info(msg string, fields ...any)
	Infof(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Warnf(msg string, fields ...any)
	Error(msg string, fields ...any)
	Errorf(msg string, fields ...any)
	Fatal(msg string, fields ...any)
	Fatalf(msg string, fields ...any)
	Panic(msg string, fields ...any)
	Panicf(msg string, fields ...any)
}

type zapLogger struct {
	logger *zap.SugaredLogger
}

func (z *zapLogger) Debug(msg string, fields ...any) {
	z.logger.Debug(msg, fields)
}

func (z *zapLogger) Debugf(msg string, fields ...any) {
	z.logger.Debugf(msg, fields)
}

func (z *zapLogger) Info(msg string, fields ...any) {
	z.logger.Info(msg, fields)
}

func (z *zapLogger) Infof(msg string, fields ...any) {
	z.logger.Infof(msg, fields)
}

func (z *zapLogger) Warn(msg string, fields ...any) {
	z.logger.Warn(msg, fields)
}

func (z *zapLogger) Warnf(msg string, fields ...any) {
	z.logger.Warnf(msg, fields)
}

func (z *zapLogger) Error(msg string, fields ...any) {
	z.logger.Error(msg, fields)
}

func (z *zapLogger) Errorf(msg string, fields ...any) {
	z.logger.Errorf(msg, fields)
}

func (z *zapLogger) Fatal(msg string, fields ...any) {
	z.logger.Fatal(msg, fields)
}

func (z *zapLogger) Fatalf(msg string, fields ...any) {
	z.logger.Fatalf(msg, fields)
}

func (z *zapLogger) Panic(msg string, fields ...any) {
	z.logger.Panic(msg, fields)
}

func (z *zapLogger) Panicf(msg string, fields ...any) {
	z.logger.Panicf(msg, fields)
}

var loggerInstance Logger

func Init(logLevel string) {
	if loggerInstance != nil {
		return
	}

	var cfg zap.Config

	if logLevel == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}

	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zLogger, err := cfg.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	loggerInstance = &zapLogger{logger: zLogger.Sugar()}
}

func GetLogger() Logger {
	if loggerInstance == nil {
		Init("development")
	}
	return loggerInstance
}

func Sync() {
	if logger, ok := loggerInstance.(*zapLogger); ok && logger != nil && logger.logger != nil {
		_ = logger.logger.Sync()
	}
}
