package main

import (
	"log"

	"github.com/MaryTissen/tg_task_bot/internal/command"
	"github.com/MaryTissen/tg_task_bot/internal/task"
	"github.com/MaryTissen/tg_task_bot/internal/tasks"
	"github.com/MaryTissen/tg_task_bot/internal/user"
	"github.com/MaryTissen/tg_task_bot/internal/users"

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

	tasks := tasks.Tasks{
		TasksMap: make(map[int][]task.Task),
	}
	users := users.Users{
		UsersMap: make(map[int]user.User),
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		_, ok := users.UsersMap[update.Message.From.ID]
		if !ok {
			users.UsersMap[update.Message.From.ID] = user.User{ //метод
				UserID:         update.Message.From.ID,
				UserCurCommand: 0,
				UserNumOfTasks: 0,
			}
		}

		switch update.Message.Command() {
		//case "help":
		//helpCommand()
		case "new":
			command.NewCommand(bot, update.Message, &tasks, &users)
			//case "edit":
			//editCommand()
			//case "delete":
			//deleteCommand()
			//case "get":
			//getCommand()
			//case "list":
			//listCommand()
		default:
			command.HandleMessage(bot, update.Message, &tasks, &users)
		}
	}
}
