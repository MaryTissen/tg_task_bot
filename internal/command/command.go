package command

import (
	"fmt"

	"github.com/MaryTissen/tg_task_bot/internal/task"
	"github.com/MaryTissen/tg_task_bot/internal/tasks"
	"github.com/MaryTissen/tg_task_bot/internal/users"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func NewCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, tasks *tasks.Tasks, users *users.Users) {
	u := users.UsersMap[inputMessage.From.ID]
	u.UserCurCommand = 2.1
	users.UsersMap[inputMessage.From.ID] = u

	task := task.Task{}
	tasks.TasksMap[inputMessage.From.ID] = append(tasks.TasksMap[inputMessage.From.ID], task)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Title:\n")
	bot.Send(msg)
	fmt.Print(users.UsersMap[inputMessage.From.ID])
}

func HandleMessage(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, tasks *tasks.Tasks, users *users.Users) {
	msg := inputMessage.Text
	u := users.UsersMap[inputMessage.From.ID]

	if u.UserCurCommand == 2.1 {
		tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Title = msg
		u.UserCurCommand = 2.2
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Date:\n")
		bot.Send(msg)
		return
	}

	if users.UsersMap[inputMessage.From.ID].UserCurCommand == 2.2 {
		tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Date = msg
		u := users.UsersMap[inputMessage.From.ID]
		u.UserCurCommand = 2.3
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Text:\n")
		bot.Send(msg)
		return
	}

	if users.UsersMap[inputMessage.From.ID].UserCurCommand == 2.3 {
		tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Text = msg
		u := users.UsersMap[inputMessage.From.ID]
		u.UserCurCommand = 2.4
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Status:\n")
		bot.Send(msg)
		return
	}

	if users.UsersMap[inputMessage.From.ID].UserCurCommand == 2.4 {
		tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Status = msg
		u := users.UsersMap[inputMessage.From.ID]
		u.UserCurCommand = 0
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Task was added\n")
		bot.Send(msg)
		return
	}
}
