package inbox

import (
	"fmt"
	"testing"

	"github.com/castlele/gogtd/src/domain/models"
	"github.com/castlele/gogtd/src/domain/repository"
	"github.com/castlele/gogtd/src/utils"
)

const (
	storageFp = "./gdt"
	inboxFp   = storageFp + "/inbox.json"
)

func TestGetAll(t *testing.T) {
	cases := []struct {
		exp []models.InboxItem
	}{
		{
			exp: []models.InboxItem{},
		},
		{
			exp: []models.InboxItem{
				{
					Id:      "1",
					Message: "hello",
				},
				{
					Id:      "2",
					Message: "world",
				},
			},
		},
	}

	utils.CreateDir(storageFp)
	defer utils.Delete(storageFp)

	for i, ts := range cases {
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			save(ts.exp)
			sut := createInteractor()

			res := sut.GetAll()

			if len(ts.exp) != len(res) {
				t.Errorf("Expected: %v; got: %v", ts.exp, res)
			}
		})
	}
}

func TestAddItem(t *testing.T) {
	t.Run("GIVEN no index.json WHEN add new item THEN index.json is created with new item", func(t *testing.T) {
		utils.Delete(storageFp)
		defer utils.Delete(storageFp)
		exp := "hello world"
		sut := createInteractor()

		_, err := sut.AddItem(exp)

		if err != nil {
			t.Errorf("Got unexpected error: %v", err)
			return
		}

		if !utils.IsExists(inboxFp) {
			t.Error("AddItem doesn't create an inbox.json")
			return
		}

		items := sut.GetAll()

		if len(items) != 1 {
			t.Errorf("Invalid items amount, has to be 1, but actual is: %v", len(items))
			return
		}

		if items[0].Message != exp {
			t.Errorf("Ivalid message in saved item. Got: %v, exp: %v", items[0].Message, exp)
			return
		}
	})
}

func TestDeleteItem(t *testing.T) {
	t.Run("GIVEN no index.json WHEN add new item THEN index.json is created with new item", func(t *testing.T) {
		utils.Delete(storageFp)
		defer utils.Delete(storageFp)
		exp := "hello world"
		sut := createInteractor()
		item, _ := sut.AddItem(exp)

		deletedItem, err := sut.DeleteItem(item.Id)

		if err != nil {
			t.Errorf("Got unexpected error: %v", err)
			return
		}

		if deletedItem.Id != item.Id {
			t.Errorf("Ivalid item was deleted. Exp: %v, got: %v", item, deletedItem)
			return
		}

		if !utils.IsExists(inboxFp) {
			t.Error("AddItem doesn't create an inbox.json")
			return
		}

		items := sut.GetAll()

		if len(items) != 0 {
			t.Errorf("Invalid items amount, has to be 0, but actual is: %v", len(items))
			return
		}
	})
}

func save(items []models.InboxItem) {
	file, err := utils.CreateFile(inboxFp)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = utils.WriteJson(file, items)

	if err != nil {
		panic(err)
	}
}

func createInteractor() Inbox {
	repo := createRepo()

	return NewInboxInteractor(repo)
}

func createRepo() repository.Repo[models.InboxItem, string] {
	repo, err := repository.NewFPRepo(inboxFp, func(item models.InboxItem) string {
		return item.Id
	})

	if err != nil {
		panic(err)
	}

	return repo
}
