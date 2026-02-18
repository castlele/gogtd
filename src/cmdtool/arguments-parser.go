package cmdtool

import (
	"flag"
	"fmt"

	"github.com/castlele/gogtd/src/commands"
)

const (
	helpMessage = `Usage:
	gogtd help

Inbox:
	gogtd inbox
	gogtd add-inbox <Message>
	gogtd update-inbox <id> [-message=<message>]
	gogtd delete-inbox <id>

Clarify:
	gogtd tasks [-box=<name>] [-project=<name>] [-favourite=<boolean>] [-status=<(pending|in_progress|done)>]
	gogtd add-task
		[-box=<name>]
		[-project=<name>]
		[-tags=<tags comma separated>]
		-message=<message>/-inbox_id=<id>
		-time=<millis>
		-energy=<low|mid|high>
	gogtd update-task <id>
		[-box=<name>]
		[-project=<name>]
		[-tags=<tags comma separated>]
		[-message=<message>]
		[-time=<millis>]
		[-energy=<low|mid|high>]
	gogtd delete-task <id>
	gogtd toggle-favourite <task_id>
	gogtd set-status <task_id> <status>(pending|in_progress|done)

Projects:
	gogtd projects
	gogtd add-project <name>
	gogtd delete-project <id>
	gogtd add-step <project_id> -message=<message>`

	inboxNoMessage    = "No message passed to create an inbox item"
	inboxNoIdToDelete = "No id passed to delete an inbox item"

	addTasksNoMandatoryArguments = "One of the arguments are not provided: " +
		"-message/-inbox_id, -time, -energy"
	taskBothIdAndMessageProvided = "You provided both inbox id and message, " +
		"you can either create task from inbox item or create completely new task"
)

func ParseArguments(args []string, factory commands.CommandsFactory) commands.Command {
	if len(args) < 2 {
		return helpCommand(factory)
	}

	switch args[1] {
	case "help":
		return helpCommand(factory)

	case "inbox":
		return factory.Inbox()
	case "add-inbox":
		if len(args) < 3 {
			return factory.Error(inboxNoMessage)
		}

		message := args[2]

		if message == "" {
			return factory.Error(inboxNoMessage)
		}

		return factory.AddInbox(message)
	case "update-inbox":
		return nil
	case "delete-inbox":
		if len(args) < 3 {
			return factory.Error(inboxNoIdToDelete)
		}

		id := args[2]

		if id == "" {
			return factory.Error(inboxNoIdToDelete)
		}

		return factory.DeleteInbox(id)

	case "tasks":
		return factory.Tasks()
	case "add-task":
		if len(args) < 5 {
			return factory.Error(addTasksNoMandatoryArguments)
		}

		fs := flag.NewFlagSet("add-task", flag.ContinueOnError)

		id := fs.String("inbox_id", "", "")
		message := fs.String("message", "", "")

		time := fs.Int64("time", 0, "")
		energy := fs.String("energy", "low", "")

		fs.Parse(args[2:])

		if *message != "" && *id != "" {
			return factory.Error(taskBothIdAndMessageProvided)
		}

		fmt.Println(*id, *message, *time, *energy)

		if *id != "" {
			return factory.AddTaskFromInbox(
				*id,
				*time,
				*energy,
			)
		} else {
			return factory.AddTask(
				*message,
				*time,
				*energy,
			)
		}
	case "update-task":
		return nil
	case "delete-task":
		if len(args) < 3 {
			return factory.Error(inboxNoIdToDelete)
		}

		id := args[2]

		if id == "" {
			return factory.Error(inboxNoIdToDelete)
		}

		return factory.DeleteTask(id)
	case "toggle-favourite":
		return nil

	case "projects":
		return nil
	case "add-project":
		return nil
	case "delete-project":
		return nil
	case "add-step":
		return nil

	default:
		return helpCommand(factory)
	}
}

func helpCommand(factory commands.CommandsFactory) commands.Command {
	return factory.Help(helpMessage)
}
