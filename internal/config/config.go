package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config holds application configuration for DBs and brokers
type Config struct {
	AppName string
	Env     string

	PostgresDSN string
	MySQLDSN    string
	Clickhouse  string
	RedisAddr   string

	KafkaBrokers []string
	NatsURL      string
	RabbitURL    string

	HTTP struct {
		Addr         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
}

// Load reads configuration from environment or config files using viper
func Load() (*Config, error) {
	v := viper.New()
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	v.SetDefault("HTTP_ADDR", ":8080")
	v.SetDefault("HTTP_READTIMEOUT", 5)
	v.SetDefault("HTTP_WRITETIMEOUT", 10)

	cfg := &Config{}
	cfg.AppName = v.GetString("APP_NAME")
	cfg.Env = v.GetString("ENV")
	cfg.PostgresDSN = v.GetString("POSTGRES_DSN")
	cfg.MySQLDSN = v.GetString("MYSQL_DSN")
	cfg.Clickhouse = v.GetString("CLICKHOUSE_DSN")
	cfg.RedisAddr = v.GetString("REDIS_ADDR")
	cfg.KafkaBrokers = v.GetStringSlice("KAFKA_BROKERS")
	cfg.NatsURL = v.GetString("NATS_URL")
	cfg.RabbitURL = v.GetString("RABBIT_URL")
	cfg.HTTP.Addr = v.GetString("HTTP_ADDR")
	cfg.HTTP.ReadTimeout = v.GetDuration("HTTP_READTIMEOUT") * time.Second
	cfg.HTTP.WriteTimeout = v.GetDuration("HTTP_WRITETIMEOUT") * time.Second

	return cfg, nil
}
