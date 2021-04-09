package main

import (
	"fmt"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

// BotConfiguration contains the basic parameters to configure the bot
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

	webhook := tb.Webhook{
		Listen:         "0.0.0.0:" + configuration.WebhookPort,
		MaxConnections: 40,
		HasCustomCert:  false,
		Endpoint: &tb.WebhookEndpoint{
			PublicURL: configuration.Webhook,
		},
	}

	fmt.Println(configuration)
	bot, err := tb.NewBot(tb.Settings{
		Token:   configuration.Token,
		Verbose: configuration.Debug,
		Poller:  &webhook,
		Reporter: func(e error) {
			fmt.Printf("%+v\n", e)
		},
	})
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, "This bot delivers you every day the Nasa Astronomic Picture Of the Day")
		addUser(m.Sender)
	})

	bot.Handle("/sendallusersthepicture", func(m *tb.Message) {
		users := getData()
		for _, user := range users {
			bot.Send(&user, "HELLO")
		}
	})

	bot.Start()
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
