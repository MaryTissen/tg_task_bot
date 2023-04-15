package main

import (
	"log"

	"github.com/MaryTissen/tg_task_bot/internal/command"
	"github.com/MaryTissen/tg_task_bot/internal/edit"
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
	tasks_edit := edit.Edit{
		EditMap: make(map[int]int),
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		_, ok := users.UsersMap[update.Message.From.ID]
		if !ok {
			users.UsersMap[update.Message.From.ID] = user.User{
				UserID:         update.Message.From.ID,
				UserCurCommand: 0,
				UserNumOfTasks: 0,
			}
		}

		switch update.Message.Command() {
		case "help":
			command.HelpCommand(bot, update.Message) //key command = 1
		case "new":
			command.NewCommand(bot, update.Message, &tasks, &users) //key command = 2
		case "edit":
			command.EditCommand(bot, update.Message, &tasks, &users) //key command = 3
			//case "delete":
			//deleteCommand()
		case "get":
			command.GetCommand(bot, update.Message, &tasks, &users) //key command = 5
		case "list":
			command.ListCommand(bot, update.Message, &tasks, &users) //key command = 6
		default:
			command.HandleMessage(bot, update.Message, &tasks, &users, &tasks_edit)
		}
	}
}
