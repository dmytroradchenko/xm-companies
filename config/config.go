package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbPort        string `env:"DB_PORT" env-description:"DB port"`
	DbName        string `env:"DB_NAME" env-description:"DB name"`
	DbHost        string `env:"DB_HOST" env-description:"DB host" env-default:"localhost"`
	DbDriverName  string `env:"DB_DRIVER_NAME" env-description:"DB driver name" env-default:"postgres"`
	DbUser        string `env:"DB_USER" env-description:"DB username"`
	DbPassword    string `env:"DB_PASSWORD" env-description:"DB user password"`
	DbSslMode     string `env:"DB_SSLMODE" env-description:"DB SSL Mode" env-default:"enable"`
	Port          string `env:"PORT" env-description:"server port"`
	TokenLifetime uint   `env:"TOKEN_LIFE_TIME" env-description:"server port" env-default:"60"`
}

func NewConfigProvider() *Config {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		panic("Failed to load configuration file: " + err.Error())
	}
	return &cfg
}
