package command

import (
	"strconv"

	"github.com/MaryTissen/tg_task_bot/internal/edit"
	"github.com/MaryTissen/tg_task_bot/internal/task"
	"github.com/MaryTissen/tg_task_bot/internal/tasks"
	"github.com/MaryTissen/tg_task_bot/internal/users"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HelpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - list all commands\n"+"/new - add new task\n"+"/edit - edit task\n"+"/delete - delete task\n"+"/get - show task by number\n"+"/list - list all tasks\n")
	bot.Send(msg)
}

func NewCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, tasks *tasks.Tasks, users *users.Users) {
	u := users.UsersMap[inputMessage.From.ID]
	u.UserCurCommand = 2.1
	users.UsersMap[inputMessage.From.ID] = u

	if u.UserNumOfTasks == len(tasks.TasksMap[inputMessage.From.ID]) {
		task := task.Task{}
		tasks.TasksMap[inputMessage.From.ID] = append(tasks.TasksMap[inputMessage.From.ID], task)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Title:\n")
	bot.Send(msg)
}

func EditCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, tasks *tasks.Tasks, users *users.Users) {
	u := users.UsersMap[inputMessage.From.ID]
	u.UserCurCommand = 3
	users.UsersMap[inputMessage.From.ID] = u

	l := strconv.Itoa(u.UserNumOfTasks)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You have saved "+l+" tasks\n"+"Type the number of task you want to edit:\n")
	bot.Send(msg)
}

func GetCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, tasks *tasks.Tasks, users *users.Users) {
	u := users.UsersMap[inputMessage.From.ID]
	u.UserCurCommand = 5
	users.UsersMap[inputMessage.From.ID] = u

	l := strconv.Itoa(u.UserNumOfTasks)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You have saved "+l+" tasks\n"+"Type the number of task you want to see:\n")
	bot.Send(msg)
}

func ListCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, tasks *tasks.Tasks, users *users.Users) {
	u := users.UsersMap[inputMessage.From.ID]
	u.UserCurCommand = 6
	users.UsersMap[inputMessage.From.ID] = u

	l := u.UserNumOfTasks - 1

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
	users.UsersMap[inputMessage.From.ID] = u
}

func HandleMessage(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, tasks *tasks.Tasks, users *users.Users, tasks_edit *edit.Edit) {
	msg := inputMessage.Text
	u := users.UsersMap[inputMessage.From.ID]

	if u.UserCurCommand == 2.1 {
		tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Number = len(tasks.TasksMap[inputMessage.From.ID])
		tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Title = msg
		u.UserCurCommand = 2.2
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Date:\n")
		bot.Send(msg)
		return
	}

	if u.UserCurCommand == 2.2 {
		tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Date = msg
		u.UserCurCommand = 2.3
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Text:\n")
		bot.Send(msg)
		return
	}

	if u.UserCurCommand == 2.3 {
		tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Text = msg
		u.UserCurCommand = 2.4
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Status (you have 2 options: done or undone):\n")
		bot.Send(msg)
		return
	}

	if u.UserCurCommand == 2.4 {
		if msg == "done" {
			tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Status = 1
		} else if msg == "undone" {
			tasks.TasksMap[inputMessage.From.ID][len(tasks.TasksMap[inputMessage.From.ID])-1].Status = 0
		} else {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid status\nTry again:\n")
			bot.Send(msg)
			return
		}

		u.UserCurCommand = 0
		u.UserNumOfTasks += 1
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Task was added\n")
		bot.Send(msg)
		return
	}

	if u.UserCurCommand == 3 {
		task_num, err := strconv.Atoi(msg)
		if err != nil {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid number\nTry again:\n")
			bot.Send(msg)
			return
		}

		tasks_edit.EditMap[inputMessage.From.ID] = task_num

		u.UserCurCommand = 3.1
		users.UsersMap[inputMessage.From.ID] = u

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Type the field you want to edit (you have 4 options: title, date, text or status):\n")
		bot.Send(msg)
		return
	}

	if u.UserCurCommand == 3.1 {
		if msg == "title" {
			u.UserCurCommand = 3.2
			users.UsersMap[inputMessage.From.ID] = u

			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "New title:\n")
			bot.Send(msg)
			return
		}
		if msg == "date" {
			u.UserCurCommand = 3.3
			users.UsersMap[inputMessage.From.ID] = u

			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "New date:\n")
			bot.Send(msg)
			return
		}
		if msg == "text" {
			u.UserCurCommand = 3.4
			users.UsersMap[inputMessage.From.ID] = u

			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "New text:\n")
			bot.Send(msg)
			return
		}
		if msg == "status" {
			u.UserCurCommand = 3.5
			users.UsersMap[inputMessage.From.ID] = u

			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "New status:\n")
			bot.Send(msg)
			return
		}
	}

	if u.UserCurCommand == 3.2 {
		task_num := tasks_edit.EditMap[inputMessage.From.ID]
		tasks.TasksMap[inputMessage.From.ID][task_num-1].Title = msg
		delete(tasks_edit.EditMap, inputMessage.From.ID)

		u.UserCurCommand = 0
		users.UsersMap[inputMessage.From.ID] = u
		return
	}
	if u.UserCurCommand == 3.3 {
		task_num := tasks_edit.EditMap[inputMessage.From.ID]
		tasks.TasksMap[inputMessage.From.ID][task_num-1].Date = msg
		delete(tasks_edit.EditMap, inputMessage.From.ID)

		u.UserCurCommand = 0
		users.UsersMap[inputMessage.From.ID] = u
		return
	}
	if u.UserCurCommand == 3.4 {
		task_num := tasks_edit.EditMap[inputMessage.From.ID]
		tasks.TasksMap[inputMessage.From.ID][task_num-1].Text = msg
		delete(tasks_edit.EditMap, inputMessage.From.ID)

		u.UserCurCommand = 0
		users.UsersMap[inputMessage.From.ID] = u
		return
	}
	if u.UserCurCommand == 3.5 {
		task_num := tasks_edit.EditMap[inputMessage.From.ID]

		if msg == "done" {
			tasks.TasksMap[inputMessage.From.ID][task_num-1].Status = 1
		} else if msg == "undone" {
			tasks.TasksMap[inputMessage.From.ID][task_num-1].Status = 0
		} else {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid status\nTry again:\n")
			bot.Send(msg)
			return
		}

		delete(tasks_edit.EditMap, inputMessage.From.ID)

		u.UserCurCommand = 0
		users.UsersMap[inputMessage.From.ID] = u
		return
	}

	if u.UserCurCommand == 5 {
		num, err := strconv.Atoi(msg)
		if err != nil {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid number\nTry again:\n")
			bot.Send(msg)
			return
		}

		l := len(tasks.TasksMap[inputMessage.From.ID]) - 1
		for i := 0; i <= l; i++ {
			if tasks.TasksMap[inputMessage.From.ID][i].Number == num {
				if tasks.TasksMap[inputMessage.From.ID][i].Status == 1 {
					msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Title: "+tasks.TasksMap[inputMessage.From.ID][i].Title+"\n"+"Date: "+tasks.TasksMap[inputMessage.From.ID][i].Date+"\n"+"Text: "+tasks.TasksMap[inputMessage.From.ID][i].Text+"\n"+"Status: "+"done"+"\n\n")
					bot.Send(msg)
				}
				if tasks.TasksMap[inputMessage.From.ID][i].Status == 0 {
					msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Title: "+tasks.TasksMap[inputMessage.From.ID][i].Title+"\n"+"Date: "+tasks.TasksMap[inputMessage.From.ID][i].Date+"\n"+"Text: "+tasks.TasksMap[inputMessage.From.ID][i].Text+"\n"+"Status: "+"undone"+"\n\n")
					bot.Send(msg)
				}
			}
		}
		u.UserCurCommand = 0
		users.UsersMap[inputMessage.From.ID] = u
		return
	}
}
