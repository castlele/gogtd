package clarify

import "github.com/castlele/gogtd/src/domain/models"

type Clarify interface {
	GetAll() []models.Task

	ConvertToTask(
		inboxItemId string,
		time int64,
		energy models.Energy,
		parent *models.TaskParent,
	) (*models.Task, error)
	AddTask(
		message string,
		time int64,
		energy models.Energy,
		parent *models.TaskParent,
	) (*models.Task, error)

	DeleteTask(id string) (*models.Task, error)

	ToggleFavourite(id string) (*models.Task, error)
	SetStatus(id string, status models.TaskStatus) (*models.Task, error)
}
