package config

type Container struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
}

type PawConfig struct {
	Name       string      `yaml:"name"`
	Containers []Container `yaml:"containers"`
	Label      string      `yaml:"label"` // checked docs labels are how u organize, label matches with cat label and becomes like nodes pod/child.
}

type CatConfig struct {
	Name  string  `yaml:"name"`
	Cpus  float64 `yaml:"cpus"`
	Ram   int64   `yaml:"ram"`
	Label string  `yaml:"label"` // dunno how to implement  it yet tho :3
}

type ClusterConfig struct {
	PawConfigs  []PawConfig `yaml:"paw"`
	CatsConfigs []CatConfig `yaml:"cat"`
}

func Parse() {}
