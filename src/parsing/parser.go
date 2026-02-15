package parsing

import (
	"github.com/castlele/gogtd/src/commands"
)

const helpMessage = `
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

func ParseArguments(args []string) commands.Command {
	if len(args) < 2 {
		return helpCommand()
	}

	switch args[1] {
	case "help":
		return helpCommand()

	case "inbox":
		return commands.NewInboxCommand()
	case "add-inbox":
		return nil
	case "update-inbox":
		return nil
	case "delete-inbox":
		return nil

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
		return helpCommand()
	}
}

func helpCommand() commands.Command {
	return commands.NewHelpCommand(helpMessage)
}
