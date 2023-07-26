package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	logger *zap.Logger
}
type Log struct {
	gorm.Model
	Level   string
	Message string
}

// NewGormLogger creates a new GormLogger with the given zap logger
func newGormLogger(logger *zap.Logger) *GormLogger {
	return &GormLogger{
		logger: logger,
	}
}

// LogMode implements the logger.Interface interface
func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	// No-op, use zap's log level instead
	return g
}

// Info implements the logger.Interface interface
func (g *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	g.logger.Sugar().Infof(msg, data...)
}

// Warn implements the logger.Interface interface
func (g *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	g.logger.Sugar().Warnf(msg, data...)
}

// Error implements the logger.Interface interface
func (g *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	g.logger.Sugar().Errorf(msg, data...)
}

// Trace implements the logger.Interface interface
func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && g.logger.Core().Enabled(zap.ErrorLevel):
		sql, rows := fc()
		if rows == -1 {
			g.logger.Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.String("sql", sql))
		} else {
			g.logger.Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.String("sql", sql), zap.Int64("rows", rows))
		}
	case elapsed > time.Second && g.logger.Core().Enabled(zap.WarnLevel):
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", time.Second)
		if rows == -1 {
			g.logger.Warn(slowLog, zap.Duration("elapsed", elapsed), zap.String("sql", sql))
		} else {
			g.logger.Warn(slowLog, zap.Duration("elapsed", elapsed), zap.String("sql", sql), zap.Int64("rows", rows))
		}
	case g.logger.Core().Enabled(zap.InfoLevel):
		sql, rows := fc()
		if rows == -1 {
			g.logger.Info("trace", zap.Duration("elapsed", elapsed), zap.String("sql", sql))
		} else {
			g.logger.Info("trace", zap.Duration("elapsed", elapsed), zap.String("sql", sql), zap.Int64("rows", rows))
		}
	}
}
