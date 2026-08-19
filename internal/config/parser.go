package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Container struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
}

type PawConfig struct {
	Name         string      `yaml:"name"`
	Containers   []Container `yaml:"containers"`
	NodeName     string      `yaml:"nodeName"`     // checked docs labels are how you organize, the label matches with cat label and becomes like node pod/child.
	DesiredState int64       `yaml:"desiredState"` // how many pods
	CatSurge     int64       `yaml:"catSurge"`     // max over desired state temp
}

type CatConfig struct {
	Name string  `yaml:"name"` // i wanna remove this for auto generated names, refer to markdown on how i decieded to do this
	Cpus float64 `yaml:"cpus"`
	Ram  int64   `yaml:"ram"`
}

type ClusterConfig struct {
	PawConfigs  []PawConfig `yaml:"paw"`
	CatsConfigs []CatConfig `yaml:"cat"`
	// add concurency part in ymal
}

func Parse(filePath string) (*ClusterConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var config ClusterConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed parse ymal content %w", filePath, err)
	}

	return &config, nil
}
