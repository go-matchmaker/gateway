package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
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
	TTL               time.Duration `yaml:"ttl"`
	ServerReadTimeout int           `yaml:"server_read_timeout"`
}

type Container struct {
	App      AppConfig      `yaml:"app"`
	Log      LogConfig      `yaml:"log"`
	HTTP     HTTPConfig     `yaml:"http"`
	Settings SettingsConfig `yaml:"settings"`
}

func LoadConfig() (*Container, error) {
	var cfg Container

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		fmt.Println("config error: ", err)
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	fmt.Println("Config loaded successfully")

	return &cfg, nil
}
