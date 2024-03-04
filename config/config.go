package config

import (
	env "github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type envConfig struct {
	DatabaseUser     string `env:"DB_USER,required"`
	DatabasePassword string `env:"DB_PASSWORD,required"`
	DatabaseHost     string `env:"DB_HOST,required"`
	DatabaseName     string `env:"DB_NAME,required"`
	DatabaseSSL      string `env:"DB_SSL,required"`
	ServerPort       string `env:"SERVER_PORT,required"`
	ServerHost       string `env:"SERVER_HOST,required"`
}

type Config struct {
	Database *DatabaseConfig
	Server   *ServerConfig
	Logger   *logrus.Logger
}

type DatabaseConfig struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabaseName     string
	DatabaseSSL      string
	DatabaseConn     *sqlx.DB
}

type ServerConfig struct {
	Port string
	Host string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	environment := envConfig{}
	err = env.Parse(&environment)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Database: &DatabaseConfig{
			DatabaseUser:     environment.DatabaseUser,
			DatabasePassword: environment.DatabasePassword,
			DatabaseHost:     environment.DatabaseHost,
			DatabaseName:     environment.DatabaseName,
			DatabaseSSL:      environment.DatabaseSSL,
		},

		Server: &ServerConfig{
			Port: environment.ServerPort,
			Host: environment.ServerHost,
		},
	}

	return cfg, nil
}

func (cfg *Config) AssignDatabaseConn(db *sqlx.DB) {
	cfg.Database.DatabaseConn = db
}

func (cfg *Config) AssignLogger(logger *logrus.Logger) {
	cfg.Logger = logger
}
