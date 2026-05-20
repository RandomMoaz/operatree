// internal/config/config.go
package config

import "fmt"

type Config struct {
	StandardDir string    `yaml:"standardDir"` // default project base dir
	Default     Project   `yaml:"default"`     // default project - used when current dir doen't include project yaml of -d is not used
	Projects    []Project `yaml:"projects"`    // tracked projects
	Daemon      Daemon    `yaml:"daemon"`      // future daemon config
}

type Project struct {
	Name     string `yaml:"name"`
	AbsPath  string `yaml:"absPath"`
	Template string `yaml:"template"` // e.g. "dev", "research"
}

type Daemon struct {
	Enabled  bool   `yaml:"enabled"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBDriver string `yaml:"dbDriver"` // sqlite, mysql, etc
	DSN      string `yaml:"dsn"`      // connection string
}

func (c *Config) ListProjects() {

	for _, p := range c.Projects {
		fmt.Println(p)
	}
}
