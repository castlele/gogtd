package cmdtool

import (
	"flag"

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
		[-parent="id::<box|project|step>"]
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
	gogtd set-status <task_id> <pending|in_progress|done>

Projects:
	gogtd projects
	gogtd add-project <name>
	gogtd delete-project <id>
	gogtd add-step <project_id> -message=<message>`

	inboxNoMessage    = "No message passed to create an inbox item"
	inboxNoIdToPassed = "No id passed to identify an inbox item"

	taskNoIdPassed = "No id passed to identify a task"

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
		command, message := parseId(factory, args, inboxNoMessage)

		if command != nil {
			return command
		}

		return factory.AddInbox(message)
	case "update-inbox":
		return nil
	case "delete-inbox":
		command, id := parseId(factory, args, inboxNoIdToPassed)

		if command != nil {
			return command
		}

		return factory.DeleteInbox(id)

	case "tasks":
		return factory.Tasks()
	case "add-task":
		return createAddTaskCommand(factory, args)
	case "update-task":
		return nil
	case "delete-task":
		command, id := parseId(factory, args, taskNoIdPassed)

		if command != nil {
			return command
		}

		return factory.DeleteTask(id)
	case "toggle-favourite":
		command, id := parseId(factory, args, taskNoIdPassed)

		if command != nil {
			return command
		}

		return factory.ToggleFavourite(id)
	case "set-status":
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

func parseId(
	factory commands.CommandsFactory,
	args []string,
	errMsg string,
) (commands.Command, string) {
	if len(args) < 3 {
		return factory.Error(errMsg), ""
	}

	id := args[2]

	if id == "" {
		return factory.Error(errMsg), ""
	}

	return nil, id
}

func createAddTaskCommand(
	factory commands.CommandsFactory,
	args []string,
) commands.Command {
	if len(args) < 5 {
		return factory.Error(addTasksNoMandatoryArguments)
	}

	fs := flag.NewFlagSet("add-task", flag.ContinueOnError)

	id := fs.String("inbox_id", "", "")
	message := fs.String("message", "", "")

	time := fs.Int64("time", 0, "")
	energy := fs.String("energy", "low", "")

	parent := fs.String("parent", "", "")

	fs.Parse(args[2:])

	if *message != "" && *id != "" {
		return factory.Error(taskBothIdAndMessageProvided)
	}

	if *id != "" {
		return factory.AddTaskFromInbox(
			*id,
			*time,
			*energy,
			*parent,
		)
	} else {
		return factory.AddTask(
			*message,
			*time,
			*energy,
			*parent,
		)
	}
}
