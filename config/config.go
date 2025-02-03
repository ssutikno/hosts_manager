package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// getConfig returns the configuration object
func GetConfig() *Config {
	return config
}
