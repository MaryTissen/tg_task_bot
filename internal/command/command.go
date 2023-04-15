package command

import (
	"fmt"
	"strconv"

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

func GetCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, tasks *tasks.Tasks, users *users.Users) {
	u := users.UsersMap[inputMessage.From.ID]
	u.UserCurCommand = 5
	users.UsersMap[inputMessage.From.ID] = u

	l := len(tasks.TasksMap[inputMessage.From.ID]) - 1

	for i := 0; i <= l; i++ {
		num := strconv.Itoa(i + 1)
		if tasks.TasksMap[inputMessage.From.ID][i].Status == 1 {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Task "+num+"\n"+"Title: "+tasks.TasksMap[inputMessage.From.ID][i].Title+"\n"+"Date: "+tasks.TasksMap[inputMessage.From.ID][i].Date+"\n"+"Text: "+tasks.TasksMap[inputMessage.From.ID][i].Text+"\n"+"Status: "+"done"+"\n\n")
			bot.Send(msg)
		}
		if tasks.TasksMap[inputMessage.From.ID][i].Status == 0 {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Task "+num+"\n"+"Title: "+tasks.TasksMap[inputMessage.From.ID][i].Title+"\n"+"Date: "+tasks.TasksMap[inputMessage.From.ID][i].Date+"\n"+"Text: "+tasks.TasksMap[inputMessage.From.ID][i].Text+"\n"+"Status: "+"undone"+"\n\n")
			bot.Send(msg)
		}
	}
	u.UserCurCommand = 0
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

	if u.UserCurCommand == 2.2 {
		tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Date = msg
		u := users.UsersMap[inputMessage.From.ID]
		u.UserCurCommand = 2.3
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Text:\n")
		bot.Send(msg)
		return
	}

	if u.UserCurCommand == 2.3 {
		tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Text = msg
		u := users.UsersMap[inputMessage.From.ID]
		u.UserCurCommand = 2.4
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Status (you have two options: done or undone):\n")
		bot.Send(msg)
		return
	}

	if u.UserCurCommand == 2.4 {
		if msg == "done" {
			tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Status = 1
		} else if msg == "undone" {
			tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Status = 0
		} else {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong status\nTry again\n")
			bot.Send(msg)
			return
		}

		u := users.UsersMap[inputMessage.From.ID]
		u.UserCurCommand = 0
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Task was added\n")
		bot.Send(msg)
		return
	}
}
