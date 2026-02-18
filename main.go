package main

import (
	"fmt"
	"os"

	"github.com/castlele/gogtd/src/cmdtool"
	"github.com/castlele/gogtd/src/commands"
	"github.com/castlele/gogtd/src/config"
	"github.com/castlele/gogtd/src/domain/clarify"
	"github.com/castlele/gogtd/src/domain/inbox"
	"github.com/castlele/gogtd/src/domain/models"
	"github.com/castlele/gogtd/src/domain/repository"
)

func main() {
	conf, err := config.LoadConfig(config.DefaultConfigFilePath)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	inboxRepo, err := repository.NewFPRepo(
		conf.GetInboxPath(),
		func(inbox models.InboxItem) string {
			return inbox.Id
		},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	tasksRepo, err := repository.NewFPRepo(
		conf.GetTasksPath(),
		func(task models.Task) string {
			return task.Id
		},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	doneTasksRepo, err := repository.NewFPRepo(
		conf.GetDoneTasksPath(),
		func(task models.Task) string {
			return task.Id
		},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	clarifyInteractor := clarify.NewClarifyInteractor(
		tasksRepo,
		doneTasksRepo,
		inboxRepo,
	)
	inboxInteractor := inbox.NewInboxInteractor(inboxRepo)

	factory := commands.NewCommandsFactory(
		inboxInteractor,
		clarifyInteractor,
		os.Stdout,
		os.Stderr,
	)
	cmd := cmdtool.ParseArguments(os.Args, factory)

	if cmd != nil {
		os.Exit(cmd.Execute())
	}

	fmt.Fprintln(os.Stderr, "invalid command usage")
	os.Exit(-1)
}
