package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5906294458:AAErtRYQ6OprFjfRt0Kvcw1UhIuF7n14xT0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand()
		case "new":
			newCommand()
		case "edit":
			editCommand()
		case "delete":
			deleteCommand()
		case "get":
			getCommand()
		case "list":
			listCommand()
		}
	}
}
