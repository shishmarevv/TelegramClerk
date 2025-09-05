package config

type Config struct {
	TelegramToken string
	Port          int
	RabbitMQURL   string
	HMACSecret    string
}
