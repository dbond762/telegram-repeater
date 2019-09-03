package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, I`m Povtoryuha bot"))
}

func main() {
	// This is need for heroku
	http.HandleFunc("/", handler)
	go http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Bot name is %s", bot.Self.UserName)

	updates := bot.ListenForWebhook("/"+bot.Token)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		answer := fmt.Sprintf("%s\n\nThis is wery useful bot", text)

		message := tgbotapi.NewMessage(chatID, answer)
		bot.Send(message)
	}
}
