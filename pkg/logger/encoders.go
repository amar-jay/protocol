package logger

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"time"
)

type RedisEncoder struct {
	zapcore.Encoder
	client *redis.Client
	key    string
	env    zapcore.EncoderConfig
}

func NewRedisEncoder(client *redis.Client, env zapcore.EncoderConfig, key string) zapcore.Encoder {
	encoder := zapcore.NewJSONEncoder(env)
	return &RedisEncoder{encoder, client, key, env}
}

func (re *RedisEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// Create a buffer to write the log message
	buf, err := re.Encoder.EncodeEntry(entry, fields)
	// Use the Redis client to store the log in Redis.
	if err != nil {
		return nil, err
	}
	_, err = re.client.LPush(context.Background(), re.key, buf, 0).Result()
	if err != nil {
		return nil, err
	}
	return buf, nil
}

type GormEncoder struct {
	zapcore.Encoder
	db *gorm.DB
}

type LogModel struct {
	gorm.Model
	Timestamp time.Time `gorm:"index"`
	Level     string
	Message   string
	Data      map[string]interface{}
}

func NewGormEncoder(db *gorm.DB) *GormEncoder {
	return &GormEncoder{db: db}
}

func (e *GormEncoder) EncodeEntry(entry zapcore.Entry, fields []zap.Field) (*buffer.Buffer, error) {
	logData := make(map[string]interface{})
	ctx := context.Background()
	defer ctx.Done()

	// Extract log fields and values
	for _, field := range fields {
		logData[field.Key] = field.String
	}

	// Create a new log model and save it to the database
	logModel := &LogModel{
		Timestamp: entry.Time,
		Level:     entry.Level.String(),
		Message:   entry.Message,
		Data:      logData,
	}

	if err := e.db.Create(logModel).Error; err != nil {
		e.db.Logger.Error(ctx, "Failed to save log to database", zap.Error(err))
		return e.Encoder.EncodeEntry(entry, fields)
	}

	return e.Encoder.EncodeEntry(entry, fields)
}
