package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Bot name is %s", bot.Self.UserName)

	updatesConfig := tgbotapi.NewUpdate(0)
	updatesConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updatesConfig)
	if err != nil {
		log.Fatal(err)
	}

	for {
		update := <-updates

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		answer := fmt.Sprintf("%s\n\nMessage send from this bot", text)

		message := tgbotapi.NewMessage(chatID, answer)
		bot.Send(message)
	}
}
