package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	godotenv "github.com/joho/godotenv"
)

// BotConfiguration an struct containing the data necessary to configure
// a working bot. It's mainly used for environment variables.
type BotConfiguration struct {
	UpdatesMethod string
	Token         string
	Debug         bool
	Webhook       string
	WebhookPort   string
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	configuration := getBotConfiguration()

	bot, err := tgbotapi.NewBotAPI(configuration.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = configuration.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	if configuration.UpdatesMethod == "updates" {
		startWithUpdates(bot, configuration)
	} else {
		startWithWebhook(bot, configuration)
	}
}

func startWithUpdates(bot *tgbotapi.BotAPI, configuration BotConfiguration) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Fatalln("Error obtaining UpdatesChannel")
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func startWithWebhook(bot *tgbotapi.BotAPI, configuration BotConfiguration) {
	_, err := bot.SetWebhook(tgbotapi.NewWebhookWithCert(configuration.Webhook+"/"+configuration.Token, "cert/cert.pem"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)

	go http.ListenAndServeTLS("0.0.0.0:"+configuration.WebhookPort, "cert/cert.pem", "cert/key.pem", nil)
	for update := range updates {
		log.Printf("%+v\n", update)
	}
}

func getBotConfiguration() BotConfiguration {
	configuration := BotConfiguration{}

	configuration.UpdatesMethod = os.Getenv("NASA_APOD_TELEGRAM_BOT_UPDATE_METHOD")
	configuration.Token = os.Getenv("NASA_APOD_TELEGRAM_BOT_TOKEN")
	configuration.Debug = getBoolFromString(os.Getenv("NASA_APOD_TELEGRAM_BOT_DEBUG"))
	configuration.Webhook = os.Getenv("NASA_APOD_TELEGRAM_BOT_WEBHOOK_URL")
	configuration.WebhookPort = os.Getenv("PORT")
	if configuration.WebhookPort == "" {
		configuration.WebhookPort = "8443"
	}

	return configuration
}

func getBoolFromString(s string) bool {
	if s == "1" || s == "true" {
		return true
	}
	return false
}
