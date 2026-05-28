package config

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"gopkg.in/yaml.v3"
)

const configFile = "operatree.yaml"

func configDir() (string, error) {

	// Respect XDG_HOME_CONFIG as first-class
	// This is a response for Github Issue
	// https://github.com/hanymamdouh82/operatree/issues/2
	xdgHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgHome != "" {
		return filepath.Join(xdgHome, "operatree"), nil
	}

	// If not available, use Go standard os package to locate config dir
	dir, err := os.UserConfigDir() // ~/.config on Linux, ~/Library/Application Support on Mac
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "operatree"), nil
}

func ConfigPath() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, configFile), nil
}

func Load() (Config, error) {
	path, err := ConfigPath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, nil // no config yet, return empty
		}
		return Config{}, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

// Saves config
func Save(cfg Config) error {
	dir, err := configDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := ConfigPath()
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// AddProject adds a project to tracked list if not already there
func AddProject(name, absPath, template string) error {
	cfg, err := Load()
	if err != nil {
		return err
	}

	for _, p := range cfg.Projects {
		if p.AbsPath == absPath {
			return nil // already tracked
		}
	}

	cfg.Projects = append(cfg.Projects, Project{
		Name:     name,
		AbsPath:  absPath,
		Template: template,
	})

	return Save(cfg)
}

// Removes project from tracked
func RemoveProject(absPath string) error {

	cfg, err := Load()
	if err != nil {
		return err
	}

	idx := slices.IndexFunc(cfg.Projects, func(p Project) bool {
		return p.AbsPath == absPath
	})

	// project is not in tracked list
	if idx == -1 {
		return fmt.Errorf("current project is not into tracked list")
	}

	nps := slices.Delete(cfg.Projects, idx, idx+1)
	cfg.Projects = nps

	return Save(cfg)
}

// SetDefaultProject sets default project to config file
func SetDefaultProject(p Project) error {
	cfg, err := Load()
	if err != nil {
		return err
	}

	cfg.Default = p
	return Save(cfg)
}
