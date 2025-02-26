package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env            string     `json:"env" env-default:"local"`
	GRPC           GRPCConfig `json:"grpc"`
	MigrationsPath string
	AuthServiceURL string
	Redis          Redis    `json:"redis"`
	Database       Database `json:"database"`
}

type GRPCConfig struct {
	Port    int           `json:"port"`
	Timeout time.Duration `json:"timeout"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if strings.TrimSpace(configPath) == "" {
		configPath = "./config/config.yaml"
	}
	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath возвращает путь к конфигурационному файлу из флага или переменной окружения.
// Приоритет: флаг > env > значение по умолчанию.
func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}

func NewConfig() (*Config, error) {
	var config Config
	if err := config.readConfig(); err != nil {
		return nil, err
	}
	return &config, nil
}

func (config *Config) readConfig() error {
	configPath := os.Getenv("CONFIG_PATH")
	if strings.TrimSpace(configPath) == "" {
		configPath = "./config/config.yaml"
	}

	fmt.Printf("Using CONFIG_PATH: %s\n", configPath)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return errors.Errorf("config file does not exist - `%s`", configPath)
	}

	err := cleanenv.ReadConfig(configPath, config)
	if err != nil {
		return errors.Wrapf(err, "failed to load config from `%s`", configPath)
	}

	fmt.Printf("Configuration loaded successfully from: %s\n", configPath)
	return nil
}