package config

import (
	"errors"

	"gopkg.in/yaml.v3"

	"github.com/Drafteame/scheduler/internal/files"
)

const defaultConfigPath = "$HOME/.scheduler.yaml"

func Load(path string) (Config, error) {
	if path == "" {
		path = defaultConfigPath
	}

	if !files.Exists(path) {
		return Config{}, errors.New("configuration file not found")
	}

	content, err := files.Read(path)
	if err != nil {
		return Config{}, errors.Join(errors.New("error reading configuration file"), err)
	}

	cfg := Config{}

	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return Config{}, errors.Join(errors.New("error unmarshalling configuration file"), err)
	}

	return cfg, nil
}

func GetJob(jobName, configPath string) (Job, bool) {
	cfg, err := Load(configPath)
	if err != nil {
		panic(err)
	}

	return cfg.GetJob(jobName)
}
