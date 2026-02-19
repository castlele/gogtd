package project

import (
	"errors"
	"testing"

	"github.com/castlele/gogtd/src/domain/models"
	"github.com/castlele/gogtd/src/domain/repository"
	"github.com/castlele/gogtd/src/utils"
	"github.com/google/uuid"
)

const (
	storageFp   = "./gdt"
	inboxFp     = storageFp + "/inbox.json"
	tasksFp     = storageFp + "/tasks.json"
	doneTasksFp = storageFp + "/done_tasks.json"
	projectsFp  = storageFp + "projects.json"
)

func TestGetAll(t *testing.T) {
	defer utils.Delete(storageFp)

	t.Run(
		"GIVEN empty repository "+
			"WHEN getting all projects "+
			"THEN returning empty projects list",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()

			projects := sut.GetAll()

			if len(projects) > 0 {
				t.Errorf(
					"Invalid amount of projects in the repo. "+
						"Expected 0, but got: %v. Tasks: %v",
					len(projects),
					projects,
				)
			}
		},
	)

	t.Run(
		"GIVEN non empty repository "+
			"WHEN getting all projects "+
			"THEN returning the same projects as in repo",
		func(t *testing.T) {
			utils.Delete(storageFp)
			project := models.Project{
				Id:   uuid.NewString(),
				Name: "Random Name",
			}
			sut := createInteractor()
			sut.projectsRepo.Create(project)

			projects := sut.GetAll()

			if len(projects) != 1 {
				t.Errorf("Invalid amount of tasks. Expected: 1, got: %v", len(projects))
				return
			}

			if projects[0].Id != project.Id {
				t.Errorf("Invalid task got. Expected: %v, got: %v", projects[0], project)
				return
			}

		},
	)
}

func TestAddProject(t *testing.T) {
	defer utils.Delete(storageFp)

	t.Run(
		"GIVEN empty projects repo "+
			"WHEN adding new project"+
			"THEN project is added",
		func(t *testing.T) {
			utils.Delete(storageFp)
			name := "work"
			sut := createInteractor()

			proj, err := sut.AddProject(name)

			if err != nil {
				panic(err)
			}

			if proj.Id == "" {
				t.Error("By adding a new project with name, Id has to be generated")
			}

			if proj.Name != name {
				t.Errorf("Name was changed during adding a new project. "+
					"Expected: %v, got: %v",
					name,
					proj.Name,
				)
			}

			if len(sut.GetAll()) == 0 {
				t.Error("Project wasn't added to the repository.")
			}
		},
	)
}

func TestDeleteProject(t *testing.T) {
	defer utils.Delete(storageFp)

	t.Run(
		"GIVEN empty projects repo "+
			"WHEN deleting none existing project "+
			"THEN nothing is changed",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()

			proj, err := sut.DeleteProject("random")

			if proj != nil {
				t.Errorf("Deleted a project, when shouldn't: %v", *proj)
			}

			if err == nil {
				t.Error("Error is nil")
			}

			if !errors.Is(err, repository.ErrNotFound) {
				t.Errorf(
					"Invalid error got. Expected: %v, got: %v",
					repository.ErrNotFound,
					err,
				)
			}
		},
	)

	t.Run(
		"GIVEN none empty projects repo "+
			"WHEN deleting existing project "+
			"THEN project is deleted from the repo",
		func(t *testing.T) {
			utils.Delete(storageFp)
			expProj := models.Project{
				Id:   "random",
				Name: "Hello_World",
			}
			sut := createInteractor()
			sut.projectsRepo.Create(expProj)

			proj, err := sut.DeleteProject(expProj.Id)

			if err != nil {
				panic(err)
			}

			if proj.Id != expProj.Id {
				t.Errorf(
					"Invalid project was deleted. Expected: %v, got: %v",
					expProj,
					*proj,
				)
			}

			if len(sut.GetAll()) != 0 {
				t.Errorf(
					"Project wasn't actually deleted from the repo. "+
						"Expected len to be 0, but got: %v",
					len(sut.GetAll()),
				)
			}
		},
	)
}

func createInteractor() *projectImpl {
	return NewProjectInteractor(createProjectsRepo())
}

func createProjectsRepo() repository.Repo[models.Project, string] {
	repo, err := repository.NewFPRepo(doneTasksFp, func(proj models.Project) string { return proj.Id })

	if err != nil {
		panic(err)
	}

	return repo
}
