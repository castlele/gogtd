package commands

import (
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/project"
)

type projectsCommand struct {
	projectsInteractor project.Project
	successOut         io.Writer
}

func newProjectsCommand(
	projectsInteractor project.Project,
	successOut io.Writer,
) *projectsCommand {
	return &projectsCommand{
		projectsInteractor: projectsInteractor,
		successOut:         successOut,
	}
}

func (this *projectsCommand) Execute() int {
	projects := this.projectsInteractor.GetAll()
	out, err := prettyPrint(projects)

	if err == nil {
		fmt.Fprintln(this.successOut, out)
	} else {
		fmt.Fprintln(this.successOut, projects)
	}

	return 0
}
