package inbox

import (
	"errors"

	"github.com/castlele/gogtd/src/domain/models"
	"github.com/castlele/gogtd/src/domain/repository"
	"github.com/google/uuid"
)

type inboxImpl struct {
	repo repository.Repo[models.InboxItem, string]
}

func NewInboxInteractor(
	repo repository.Repo[models.InboxItem, string],
) Inbox {
	return &inboxImpl{
		repo: repo,
	}
}

func (this *inboxImpl) GetAll() []models.InboxItem {
	items, err := this.repo.List()

	if err != nil {
		return make([]models.InboxItem, 0)
	}

	return items
}

func (this *inboxImpl) AddItem(message string) (*models.InboxItem, error) {
	item := createInboxItem(message)
	err := this.repo.Create(item)

	return &item, err
}

func (this *inboxImpl) DeleteItem(id string) (*models.InboxItem, error) {
	items := this.GetAll()

	if len(items) == 0 {
		return nil, errors.New("Can't delete item from empty repository")
	}

	item, err := this.repo.Delete(id)

	return &item, err
}

func createInboxItem(message string) models.InboxItem {
	return models.InboxItem{
		Id:      uuid.NewString(),
		Message: message,
	}
}
