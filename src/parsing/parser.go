package parsing

import (
	"github.com/castlele/gogtd/src/commands"
)

const (
	helpMessage = `

	Usage:
		gogtd help

	Inbox:
		gogtd inbox
		gogtd add-inbox <Message>
		gogtd update-inbox <id> [--message=<message>]
		gogtd delete-inbox <id>

	Clarify:
		gogtd tasks [--box=<name>] [--project=<name>] [--favourite=<boolean>]
		gogtd add-task
			[--box=<name>]
			[--project=<name>]
			[--tags=<tags comma separated>]
			--message=<message>
			--time=<millis>
			--energy=<low|mid|high>
		gogtd update-task <id>
			[--box=<name>]
			[--project=<name>]
			[--tags=<tags comma separated>]
			[--message=<message>]
			[--time=<millis>]
			[--energy=<low|mid|high>]
		gogtd delete-task <id>
		gogtd toggle-favourite <task_id>

	Projects:
		gogtd projects
		gogtd add-project <name>
		gogtd delete-project <id>
		gogtd add-step <project_id> --message=<message>`

	inboxNoMessage    = "No message passed to create an inbox item"
	inboxNoIdToDelete = "No id passed to delete an inbox item"
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
		return nil
	case "add-task":
		return nil
	case "update-task":
		return nil
	case "delete-task":
		return nil
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
