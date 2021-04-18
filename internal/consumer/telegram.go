package consumer

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/michalq/thebillbot/internal/messenger"
)

type TelegramConsumer struct {
	botAPIClient *tgbotapi.BotAPI
	messenger    *messenger.Messenger
}

func NewTelegramConsumer(botAPIClient *tgbotapi.BotAPI, msg *messenger.Messenger) *TelegramConsumer {
	return &TelegramConsumer{botAPIClient, msg}
}

func (t *TelegramConsumer) Listen() error {

	log.Printf("Authorized on account %s", t.botAPIClient.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := t.botAPIClient.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		t.botAPIClient.Send(msg)
		answers := t.messenger.Scan(update.Message.Text)
		if len(answers) > 0 {
			for _, answer := range answers {
				t.botAPIClient.Send(tgbotapi.NewMessage(update.Message.Chat.ID, answer))
			}
		} else {
			t.botAPIClient.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "I don't understand"))
		}
	}
	return nil
}
