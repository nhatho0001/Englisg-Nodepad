package configs

import (
	"fmt"
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Configs struct {
	DB_NAME      string `env:"DB_NAME,required"`
	DB_HOST      string `env:"DB_HOST,required"`
	DB_USER      string `env:"DB_USER,required"`
	DB_PASSWORLD string `env:"DB_PASSWORLD,required"`
	DB_PORT      string `env:"DB_PORT,required"`
	PORT         string `env:"SERVER_PORT"`
	HOST         string `env:"SERVER_HOST"`
	JWT_SECRET   string `env:"JWT_SECRET"`
}

func NewConfig() (*Configs, error) {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
		return nil, err
	}

	cfg := Configs{}
	if err := env.Parse(&cfg); err != nil {
		slog.Error("Parse config env file faild")
		return nil, err
	}

	slog.Info("Config Loaded")
	return &cfg, nil
}

func (cfg Configs) DataBaseURl() string {
	// postgres://username:password@localhost:5432/database_name
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DB_USER, cfg.DB_PASSWORLD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
}
