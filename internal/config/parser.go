package config

type Container struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
}

type PawConfig struct {
	Name         string      `yaml:"name"`
	Containers   []Container `yaml:"containers"`
	NodeName     string      `yaml:"nodeName"`     // checked docs labels are how u organize, label matches with cat label and becomes like nodes pod/child.
	DesiredState int64       `yaml:"desiredState"` // how many pods
	CatSurge     int64       `yaml:"catSurge"`     // max over desired state temp
}

type CatConfig struct {
	Name string  `yaml:"name"`
	Cpus float64 `yaml:"cpus"`
	Ram  int64   `yaml:"ram"`
}

type ClusterConfig struct {
	PawConfigs  []PawConfig `yaml:"paw"`
	CatsConfigs []CatConfig `yaml:"cat"`
}

func Parse() {}
