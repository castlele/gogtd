package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/castlele/gogtd/src/domain/models"
)

func prettyPrint(obj any) (string, error) {
	bytes, err := json.MarshalIndent(obj, "", "  ")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func parseStatus(status string) (models.TaskStatus, error) {
	if status == "" {
		return models.TaskStatusPending, errors.New("You provided empty status")
	}

	switch status {
	case "pending":
		return models.TaskStatusPending, nil
	case "in_progress":
		return models.TaskStatusInProgress, nil
	case "done":
		return models.TaskStatusDone, nil
	default:
		return models.TaskStatusPending, fmt.Errorf(
			"Invalid status provided. Expected one of: pending, in_progress, done. Got: %v",
			status,
		)
	}
}
