package commands

import (
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/project"
)

type deleteProjectCommand struct {
	projectsInteractor project.Project
	successOut         io.Writer
	errOut             io.Writer
	id                 string
}

func newDeleteProjectCommand(
	id string,
	projectsInteractor project.Project,
	successOut io.Writer,
	errOut io.Writer,
) *deleteProjectCommand {
	return &deleteProjectCommand{
		projectsInteractor: projectsInteractor,
		successOut:         successOut,
		errOut:             errOut,
		id:                 id,
	}
}

func (this *deleteProjectCommand) Execute() int {
	proj, err := this.projectsInteractor.DeleteProject(this.id)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	out, err := prettyPrint(proj)

	if err != nil {
		fmt.Fprintln(this.successOut, proj)
	} else {
		fmt.Fprintln(this.successOut, out)
	}

	return 0
}
