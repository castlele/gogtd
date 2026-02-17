package inbox

import (
	"github.com/castlele/gogtd/src/domain/models"
)

type Inbox interface {
	GetAll() []models.InboxItem
	AddItem(message string) (*models.InboxItem, error)
	DeleteItem(id string) (*models.InboxItem, error)
}
