// internal/config/config.go
package config

type Config struct {
	StandardDir string    `yaml:"standardDir"` // default project base dir
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
