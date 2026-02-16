package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/castlele/gogtd/src/utils"
)

const DefaultConfigFilePath = ".config/gogtd/config.json"

func LoadConfig(path string) (*Config, error) {
	home, _ := os.UserHomeDir()
	path = filepath.Join(home, path)
	file, _ := utils.OpenFile(path)
	config := createDefaultConfig()

	if file != nil {
		defer file.Close()
		bytes, _ := utils.LoadBytesFromFile(file)
		err := json.Unmarshal(bytes, config)

		_, ok := err.(*json.SyntaxError)

		if err != nil && ok {
			return nil, err
		}
	}

	return config, nil
}

func createDefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	storagePath := filepath.Join(home, ".config/gogtd/storage/")

	return &Config{
		Storage: storagePath,
	}
}
