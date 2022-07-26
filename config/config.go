package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Mongo struct {
			Connection struct {
				Text string `yaml:"text"`
			} `yaml:"connection"`
		} `yaml:"mongo"`
	} `yaml:"database"`
}

func (config *Config) SetConfig() error {
	file, err := os.Open("config.yml")
	defer file.Close()
	if err != nil {
		return err
	}
	err = yaml.NewDecoder(file).Decode(&config)

	if err != nil {
		return err
	}
	return nil
}
