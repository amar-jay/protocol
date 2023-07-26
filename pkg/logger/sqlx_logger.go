package logger

import (
	"time"

	"go.uber.org/zap"
)

// CustomZapSQLXLogger is a custom logger that implements sqlx.Logger interface
type CustomZapSQLXLogger struct {
	logger *zap.Logger
}

// NewCustomZapSQLXLogger creates a new instance of CustomZapSQLXLogger
func newCustomZapSQLXLogger(logger *zap.Logger) *CustomZapSQLXLogger {
	return &CustomZapSQLXLogger{logger: logger}
}

// Print logs the given SQL query and its arguments
func (c *CustomZapSQLXLogger) Print(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	query := v[0].(string)
	args := v[1].([]interface{})
	duration := v[2].(time.Duration)

	c.logger.Info("SQL Query Executed",
		zap.String("query", query),
		zap.Any("args", args),
		zap.Duration("duration", duration),
	)
}
