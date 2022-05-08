package main

import (
	"log"
	"os"
	"github.com/yanzay/tbot/v2"
)

var (
	app     application
	bot     *tbot.Server
	token   string
)

type application struct {
	client *tbot.Client
}

func init() {
	token = os.Getenv("TELEGRAM_TOKEN")
}

func main() {
	bot = tbot.New(token, tbot.WithWebhook("https://reshalfahsi-elementalbot.herokuapp.com", ":"+os.Getenv("PORT")))
	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	bot.HandleMessage("/predict", app.predictHandler)
	log.Fatal(bot.Start())
}
