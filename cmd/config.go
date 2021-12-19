package cmd

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Profile map[string]Config

type Config struct {
	Domain       string `yaml:"domain"`
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
}

func LoadConfig() (Profile, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(home, ".auth0", "config.yaml")
	if _, err := os.Stat(filePath); err != nil {
		return nil, fmt.Errorf("config file %s not found", filePath)
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var profile Profile
	err = yaml.Unmarshal(b, &profile)
	return profile, err
}
