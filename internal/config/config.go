package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DB_Url   string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func (c Config) SetUser(name string) error {
	c.Username = name
	err := write(c)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig() error {
	cfg := Config{
		DB_Url:   "postgres://username@localhost:5432/gator",
		Username: "",
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	by, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := json.Unmarshal(by, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
}
