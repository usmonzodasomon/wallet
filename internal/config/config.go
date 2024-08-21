package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Database
	SecretKey string `env:"SECRET_KEY"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Database struct {
	PostgresHost     string `env:"POSTGRES_HOST" env-default:"127.0.0.1"`
	PostgresPort     string `env:"POSTGRES_PORT" env-default:"5432"`
	PostgresUser     string `env:"POSTGRES_USER" env-default:"postgres"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" env-required:"true"`
	PostgresDatabase string `env:"POSTGRES_DATABASE" env-default:"postgres"`
}

var Cfg Config

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("cannot load env: %s", err)
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if err := cleanenv.ReadConfig(configPath, &Cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &Cfg
}
