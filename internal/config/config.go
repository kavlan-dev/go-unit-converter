package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Environment string // dev, prod
	CORS        []string
}

func InitConfig() (*Config, error) {
	var config Config
	config.Environment = envOrDefault("ENV", "dev")
	config.CORS = strings.Split(envOrDefault("CORS", "*"), ",")

	if config.Environment != "dev" && config.Environment != "prod" {
		return nil, fmt.Errorf("Не верно указано окружение %s", config.Environment)
	}

	return &config, nil
}

func envOrDefault(varName string, defaultValue string) string {
	value := os.Getenv(varName)
	if value == "" {
		value = defaultValue
	}

	return value
}
