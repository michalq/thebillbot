package main

import (
	"fmt"
	"log"

	"github.com/Netflix/go-env"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/michalq/thebillbot/internal/config"
	"github.com/michalq/thebillbot/internal/consumer"
	"github.com/michalq/thebillbot/internal/crypto"
	"github.com/michalq/thebillbot/internal/messenger"
	"github.com/michalq/thebillbot/pkg/arcgis"
	"github.com/michalq/thebillbot/pkg/coinbase"
	"github.com/michalq/thebillbot/pkg/szczepimysie"

	topicCovid "github.com/michalq/thebillbot/pkg/topics/covid"
	topicCrypto "github.com/michalq/thebillbot/pkg/topics/crypto"
)

func main() {

	fmt.Println("Starting app...")
	cfg := config.AppConfig{}
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Panic(fmt.Sprintf("read config err: %+v", err))
	}
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Panic(fmt.Sprintf("telegram initialization: %+v", err))
	}
	bot.Debug = true

	provider := crypto.NewCoinbasePriceProvider(coinbase.NewClient(cfg.Coinbase.ApiKey, cfg.Coinbase.SecretKey))
	szczepimysieClient := szczepimysie.NewClient(arcgis.NewClient("zk7YlClTgerl62BY"), "global_szczepienia_widok2")

	msg := messenger.NewMessenger()
	msg.AddTopic(topicCrypto.NewConversion(provider))
	msg.AddTopic(topicCovid.NewVaccinationStatus(szczepimysieClient))
	msg.AddTopic(messenger.NewHelpTopic(msg))

	cs := consumer.NewTelegramConsumer(bot, msg)
	if err := cs.Listen(); err != nil {
		log.Panic(fmt.Sprintf("telegram listener: %+v", err))
	}
}
