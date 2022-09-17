package config

import (
	"github.com/spf13/viper"
)

type (
	AppMode string
	Config  struct {
		AppMode  AppMode
		Database DatabaseConfig
	}

	DatabaseConfig struct {
		Sqlite3 Sqlite3Config
	}

	Sqlite3Config struct {
		LocalPath string
	}
)

func Init(configsDir string) (*Config, error) {
	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("app_mode", &cfg.AppMode); err != nil {
		return err
	}

	return viper.UnmarshalKey("db.local_path", &cfg.Database.Sqlite3.LocalPath)
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}
