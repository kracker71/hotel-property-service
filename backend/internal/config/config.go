package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
		Env  string
	}
	Database struct {
		Driver    string
		DSN       string
		Migration bool
		Seeding   bool
	}
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	slog.Info("[CONFIG]", "message", "Configuration loaded successfully")
	return &cfg, nil
}
