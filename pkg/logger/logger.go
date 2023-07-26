package logger

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type LoggerConfig struct {
	Redis          *redis.Client
	Gorm           *gorm.DB
	FileLogging    bool // if true, logs will be written to a file
	DbLogging      bool // if true, logs will be written to a database
	ConsoleLogging bool `default:"true"` // if true, logs will be written to console
}

// logger is not specific to any package
func NewDefaultLogger(conf LoggerConfig) *zap.Logger {
	// select based on environment "ENV"
	e := os.Getenv("ENV")
	var envConf, prodConf, devConf zapcore.EncoderConfig
	prodConf = zap.NewProductionEncoderConfig()
	devConf = zap.NewDevelopmentEncoderConfig()
	var zapcores []zapcore.Core

	if e == "production" {
		envConf = prodConf
		envConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		envConf = devConf
	}

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	stdout := zapcore.AddSync(os.Stdout)
	level := zap.NewAtomicLevel()
	alertLevel := zap.NewAtomicLevelAt(zapcore.WarnLevel)
	if conf.FileLogging {
		zapcores = append(zapcores, zapcore.NewCore(zapcore.NewJSONEncoder(prodConf), file, level))
	}

	if !conf.ConsoleLogging {
		zapcores = append(zapcores, zapcore.NewCore(zapcore.NewConsoleEncoder(envConf), stdout, level))
	}

	if conf.DbLogging && conf.Gorm != nil {
		zapcores = append(zapcores, zapcore.NewCore(NewGormEncoder(conf.Gorm), nil, alertLevel))
	}

	if conf.Redis != nil {
		zapcores = append(zapcores, zapcore.NewCore(NewRedisEncoder(conf.Redis, prodConf, "log"), nil, level))
	}

	core := zapcore.NewTee(
		zapcores...,
	)
	logger := zap.New(core)
	return logger

}

func NewGormLogger(logger *zap.Logger) *GormLogger {

	return newGormLogger(logger)
}

func NewSQLLogger(logger *zap.Logger) *CustomZapSQLXLogger {

	return newCustomZapSQLXLogger(logger)
}
