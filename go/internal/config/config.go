package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port               string        `mapstructure:"PORT"`
	LogLevel           string        `mapstructure:"LOG_LEVEL"`
	DBHost             string        `mapstructure:"DB_HOST"`
	DBPort             string        `mapstructure:"DB_PORT"`
	DBUser             string        `mapstructure:"DB_USER"`
	DBPassword         string        `mapstructure:"DB_PASSWORD"`
	DBName             string        `mapstructure:"DB_NAME"`
	DBSSLMode          string        `mapstructure:"DB_SSLMODE"`
	DBMaxOpenConns     int           `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns     int           `mapstructure:"DB_MAX_IDLE_CONNS"`
	DBConnMaxLifetime  time.Duration `mapstructure:"DB_CONN_MAX_LIFETIME"`
	RedisHost          string        `mapstructure:"REDIS_HOST"`
	RedisPort          string        `mapstructure:"REDIS_PORT"`
	RedisPassword      string        `mapstructure:"REDIS_PASSWORD"`
	RedisDB            int           `mapstructure:"REDIS_DB"`
	RedisCacheTTL      time.Duration `mapstructure:"REDIS_CACHE_TTL"`
	RabbitMQHost       string        `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort       string        `mapstructure:"RABBITMQ_PORT"`
	RabbitMQUser       string        `mapstructure:"RABBITMQ_USER"`
	RabbitMQPassword   string        `mapstructure:"RABBITMQ_PASSWORD"`
	JWTSecret          string        `mapstructure:"JWT_SECRET"`
	JWTExpiry          time.Duration `mapstructure:"JWT_EXPIRY"`
	RateLimitRequests  int           `mapstructure:"RATE_LIMIT_REQUESTS"`
	RateLimitWindow    time.Duration `mapstructure:"RATE_LIMIT_WINDOW"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetDefault("PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("DB_NAME", "multitenant")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("DB_MAX_OPEN_CONNS", 25)
	viper.SetDefault("DB_MAX_IDLE_CONNS", 25)
	viper.SetDefault("DB_CONN_MAX_LIFETIME", "5m")
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("REDIS_CACHE_TTL", "5m")
	viper.SetDefault("RABBITMQ_HOST", "localhost")
	viper.SetDefault("RABBITMQ_PORT", "5672")
	viper.SetDefault("RABBITMQ_USER", "guest")
	viper.SetDefault("RABBITMQ_PASSWORD", "guest")
	viper.SetDefault("JWT_SECRET", "super-secret-key-change-me")
	viper.SetDefault("JWT_EXPIRY", "24h")
	viper.SetDefault("RATE_LIMIT_REQUESTS", 100)
	viper.SetDefault("RATE_LIMIT_WINDOW", "1m")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
