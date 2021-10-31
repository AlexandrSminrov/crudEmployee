package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// ServerConfig server configuration
type ServerConfig struct {
	Name           string        `envconfig:"NAME" default:"employees"`
	Version        string        `envconfig:"VERSION" required:"true"`
	Port           string        `envconfig:"PORT" required:"true"`
	MaxHeaderBytes int           `envconfig:"MAX_HEADER_BYTES" required:"true"`
	ReadTimeout    time.Duration `envconfig:"READ_TIMEOUT" required:"true"`
	WriteTimeout   time.Duration `envconfig:"WRITE_TIMEOUT" required:"true"`
	DBConfig       DBConfig
}

// DBConfig database configuration
type DBConfig struct {
	MaxConn      int           `envconfig:"DB_MAX_CONN" default:"10"`
	MaxIdleConn  int           `envconfig:"DB_MAX_IDLE_CONN" default:"10"`
	Port         int           `envconfig:"DB_PORT" required:"true"`
	Host         string        `envconfig:"DB_HOST" required:"true"`
	User         string        `envconfig:"DB_USER" required:"true"`
	Password     string        `envconfig:"DB_PASSWORD" required:"true"`
	DBName       string        `envconfig:"DB_NAME" required:"true"`
	TimeIdleConn time.Duration `envconfig:"DB_TIME_IDLE_CONN" default:"5m"`
}

// GetConfig initialize configuration
func GetConfig() (*ServerConfig, error) {
	config := &ServerConfig{}
	return config, envconfig.Process("", config)
}
