package config

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	log "github.com/fallibilism/protocol/pkg/logger"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Config
	DB    *sql.DB
	Gorm  *gorm.DB
	Redis *redis.Client // redis client
}

type Config struct {
	Name         string           `yaml:"name"`
	Developement bool             `yaml:"developement"`
	Port         uint             `yaml:"port"`
	JWTSecret    string           `yaml:"jwt_secret"` // LTI and JWT secret
	JWTIssuer    string           `yaml:"jwt_issuer"`
	ConsumerKey  string           `yaml:"consumer_key"`
	LogSettings  LogConfig        `yaml:"log_settings"`
	ServerConfig ServerConfig     `yaml:"server"`
	Openai       OpenAIConfig     `yaml:"open_ai"`
	Postgres     PostgresConfig   `yaml:"postgres"`
	Redis        RedisConfig      `yaml:"redis"`
	Livekit      LivekitConfig    `yaml:"livekit"`
	Prometheus   PrometheusConfig `yaml:"prometheus"`
}

const (
	ViewExt = ".html"
)

var (
	Prometheus = PrometheusConfig{
		Enabled:     true,
		Namespace:   "v",
		MetricsPath: "/metrics",
	}
	Server = ServerConfig{
		Port:     "8080",
		Host:     "0.0.0.0",
		ViewPath: "views",
	}
	Livekit = LivekitConfig{
		Host:   "http://localhost:7880",
		ApiKey: "api_key",
		Secret: "secret",
	}
)

type LogConfig struct {
	LogFile    string `yaml:"log_file"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}

type PrometheusConfig struct {
	Enabled bool
	// Namespace for prometheus metrics
	Namespace string
	// MetricsPath for prometheus metrics
	MetricsPath string
}

type ServerConfig struct {
	// Port for server
	Port string
	// Host for server
	Host string
	// view path for server
	ViewPath string
}

type LivekitConfig struct {
	Host   string `yaml:"host"`
	ApiKey string `yaml:"api_key"`
	Secret string `yaml:"secret"`
}

type RedisConfig struct {
	Address     string `yaml:"address"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Port        int32  `yaml:"port"`
	DBName      int    `yaml:"db"`
	UseTLS      bool   `yaml:"use_tls"`
	MaxRetries  int    `yaml:"max_retries"`
	IdleTimeout int    `yaml:"idle_timeout"` // in seconds
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"database"`
	Prefix   string `yaml:"prefix"`
	SslMode  string `yaml:"sslmode" default:"disable"`
	TimeZone string `yaml:"timezone" default:"Asia/Jakarta"`

	External bool   `yaml:"external"`
	URI      string `env:"POSTGRES_URI"`
}

// OPEN AI CONFIG
type OpenAIConfig struct {
	ApiKey string `yaml:"api_key"`
	Secret string `yaml:"secret"`
}

var logger = log.NewDefaultLogger(log.LoggerConfig{})

func SetConfig(config Config) (app *App) {

	if config.Openai.ApiKey == "" || config.Openai.Secret == "" {
		logger.Panic("config: openai api key and secret must be set")
	}

	if config.Postgres.Host == "" || config.Postgres.Username == "" || config.Postgres.Password == "" || config.Postgres.DBName == "" {
		logger.Panic("config: postgres host, username, password, and database name must be set")
	}

	// environmental variables override config
	if v := os.Getenv("PORT"); v != "" {
		port, err := strconv.Atoi(v)
		if err == nil {
			config.Port = uint(port)
		}
	}

	if v := os.Getenv("JWT_SECRET"); v != "" {
		config.JWTSecret = v
	}

	if v := os.Getenv("POSTGRES_URI"); v != "" {
		config.Postgres.URI = v
	}

	gormDB, sqlDB, err := newDBClient()
	if err != nil {
		logger.Panic("config: " + err.Error())
	}

	redis, err := newRedisClient(config.Redis)

	return &App{config, sqlDB, gormDB, redis}
}

func newDBClient() (*gorm.DB, *sql.DB, error) {
	sqlDB, err := sql.Open("pgx", "mydb_dsn")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return gormDB, sqlDB, err
}

func newRedisClient(conf RedisConfig) (*redis.Client, error) {
	if conf.Address == "" || conf.Username == "" || conf.Password == "" {
		logger.Panic("config: redis host, username, password, and database name must be set")
	}
	client := redis.NewClient(&redis.Options{
		Addr:        conf.Address,
		Username:    conf.Username,
		Password:    conf.Password,
		DB:          conf.DBName,
		MaxRetries:  conf.MaxRetries,
		IdleTimeout: time.Duration(conf.IdleTimeout) * time.Second,
	})

	return client, nil
}
