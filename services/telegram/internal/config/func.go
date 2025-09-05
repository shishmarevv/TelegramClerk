package config

import (
	"fmt"
	"os"
	"strconv"
)

func Load() (*Config, error) {
	telegramToken, err := mustGet("TELEGRAM_BOT_TOKEN")
	if err != nil {
		return nil, fmt.Errorf("error loading telegram token: %v", err)
	}

	rabbitMQURL := defaultGet("RABBITMQ_URL", "stub")

	hmacSecret, err := mustGet("HMAC_SECRET")
	if err != nil {
		return nil, fmt.Errorf("error loading hmac secret: %v", err)
	}

	portStr := defaultGet("PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("error loading port: %v", err)
	}

	return &Config{
		TelegramToken: telegramToken,
		Port:          port,
		RabbitMQURL:   rabbitMQURL,
		HMACSecret:    hmacSecret,
	}, nil
}

func mustGet(key string) (string, error) {
	value, err := os.LookupEnv(key)
	if !err || value == "" {
		return "", fmt.Errorf("environment variable %s is required", key)
	}
	return value, nil
}

func defaultGet(key, defaultValue string) string {
	value, err := os.LookupEnv(key)
	if !err || value == "" {
		return defaultValue
	}
	return value
}
