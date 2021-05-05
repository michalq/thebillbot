package config

type AppConfig struct {
	Telegram struct {
		Token string `env:"TELEGRAM_TOKEN"`
	}
	Coinbase struct {
		ApiKey    string `env:"COINBASE_API_KEY"`
		SecretKey string `env:"COINBASE_API_SECRET_KEY"`
	}
	Szczepimysie struct {
		ClientId    string `env:"SZCZEPIMYSIE_CLIENT_ID"`
		ServiceName string `env:"SZCZEPIMYSIE_SERVICE_NAME"`
	}
}
