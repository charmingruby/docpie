package config

import (
	env "github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type envConfig struct {
	DatabaseUser        string `env:"DB_USER,required"`
	DatabasePassword    string `env:"DB_PASSWORD,required"`
	DatabaseHost        string `env:"DB_HOST,required"`
	DatabaseName        string `env:"DB_NAME,required"`
	DatabaseSSL         string `env:"DB_SSL,required"`
	ServerPort          string `env:"SERVER_PORT,required"`
	ServerHost          string `env:"SERVER_HOST,required"`
	JwtSecretKey        string `env:"JWT_SECRET_KEY,required"`
	CloudflareAccountID string `env:"CLOUDFLARE_ACCOUNT_ID,required"`
	AWSBucketName       string `env:"AWS_BUCKET_NAME,required"`
	AWSAccessKeyID      string `env:"AWS_ACCESS_KEY_ID,required"`
	AWSSecretAccessKey  string `env:"AWS_SECRET_ACCESS_KEY,required"`
}

type Config struct {
	Database   *DatabaseConfig
	Server     *ServerConfig
	Cloudflare *CloudflareConfig
	Logger     *logrus.Logger
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Name     string
	SSL      string
	Conn     *sqlx.DB
}

type ServerConfig struct {
	Port string
	Host string
}

type CloudflareConfig struct {
	AccountID       string
	BucketName      string
	AccessKeyID     string
	SecretAccessKey string
}

func New(logger *logrus.Logger) (*Config, error) {
	logger.Info("Loading configuration...")

	environment := envConfig{}
	err := env.Parse(&environment)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Database: &DatabaseConfig{
			User:     environment.DatabaseUser,
			Password: environment.DatabasePassword,
			Host:     environment.DatabaseHost,
			Name:     environment.DatabaseName,
			SSL:      environment.DatabaseSSL,
		},

		Server: &ServerConfig{
			Port: environment.ServerPort,
			Host: environment.ServerHost,
		},
		Cloudflare: &CloudflareConfig{
			AccountID:       environment.CloudflareAccountID,
			BucketName:      environment.AWSBucketName,
			AccessKeyID:     environment.AWSAccessKeyID,
			SecretAccessKey: environment.AWSSecretAccessKey,
		},
		Logger: logger,
	}

	logger.Info("Configuration done.")

	return cfg, nil
}

func (cfg *Config) SetDatabaseConn(db *sqlx.DB) {
	cfg.Database.Conn = db
}
