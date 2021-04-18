package config

type AppConfig struct {
	Telegram struct {
		Token string `env:"TELEGRAM_TOKEN"`
	}
	Coinbase struct {
		ApiKey    string `env:"COINBASE_API_KEY"`
		SecretKey string `env:"COINBASE_API_SECRET_KEY"`
	}
	Binance struct {
		ApiKey    string `env:"BINANCE_API_KEY"`
		SecretKey string `env:"BINANCE_API_SECRET_KEY"`
	}
}
