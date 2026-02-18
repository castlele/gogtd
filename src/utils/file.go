package utils

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func LoadBytesFromFile(file *os.File) ([]byte, error) {
	if file == nil {
		return nil, errors.New("os.File wasn't provided for reading!")
	}

	bytes, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func OpenFile(path string) (*os.File, error) {
	return os.Open(path)
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)

	return file, err
}

func Delete(path string) error {
	return os.RemoveAll(path)
}

func WriteJson(file *os.File, obj any) error {
	bytes, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	_, err = file.Write(bytes)

	return err
}

func IsExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil && !errors.Is(err, os.ErrNotExist)
}
