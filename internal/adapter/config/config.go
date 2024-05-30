package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = "config.yml"
)

type AppConfig struct {
	Name string `yaml:"name"`
}

type LogConfig struct {
	Level int `yaml:"level"`
}

type HTTPConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type SettingsConfig struct {
	TTL               int `yaml:"ttl"`
	ServerReadTimeout int `yaml:"server_read_timeout"`
}

type ServiceConfig struct {
	URL      string `yaml:"url"`
	Cache    bool   `yaml:"cache"`
	TTL      int    `yaml:"ttl"`
	LogLevel int    `yaml:"log_level"`
}

type Config struct {
	App      AppConfig      `yaml:"app"`
	Log      LogConfig      `yaml:"log"`
	HTTP     HTTPConfig     `yaml:"http"`
	Settings SettingsConfig `yaml:"settings"`
	Services struct {
		Auth       ServiceConfig `yaml:"auth"`
		Management ServiceConfig `yaml:"management"`
	} `yaml:"services"`
}

func LoadConfig() (*Config, error) {
	var cfg *Config

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
