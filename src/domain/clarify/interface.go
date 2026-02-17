package clarify

import "github.com/castlele/gogtd/src/domain/models"

type Clarify interface {
	GetAll() []models.Task
	AddTask(inboxItemId string) (*models.Task, error)
	DeleteTask(id string) (*models.Task, error)
}
