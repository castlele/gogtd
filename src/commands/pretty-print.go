package commands

import "encoding/json"

func prettyPrint(obj any) (string, error) {
	bytes, err := json.MarshalIndent(obj, "", "  ")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
