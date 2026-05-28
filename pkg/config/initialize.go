package config

import (
	"os"

	"github.com/charmbracelet/huh"
)

func InitializeConfig() error {

	// Check if config already exists
	existing, err := Load()
	if err != nil {
		return err
	}

	var standardDir string
	var overwrite bool
	var defaultFileManager string

	if existing.StandardDir != "" {
		// Config already exists — ask before overwriting
		huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("Config already exists. Overwrite?").
					Value(&overwrite),
			),
		).Run()

		if !overwrite {
			return nil
		}
		standardDir = existing.StandardDir // pre-fill with existing value
	}

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Standard projects directory").
				Description("Default location where all projects will be created").
				Placeholder("/home/user/projects").
				Value(&standardDir),
		),
	).Run()
	if err != nil {
		return err
	}

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Default file manager").
				Description("Default binary name for file manager").
				Placeholder("yazi").
				Value(&defaultFileManager),
		),
	).Run()
	if err != nil {
		return err
	}

	// get default editor
	editor := os.Getenv("EDITOR")

	cfg := Config{
		StandardDir: standardDir,
		Editor:      editor,
		FileManager: defaultFileManager,
		Projects:    existing.Projects, // preserve tracked projects if overwriting
		Daemon: Daemon{
			Enabled:  false,
			Host:     "localhost",
			Port:     7070,
			DBDriver: "sqlite",
		},
	}

	if err := Save(cfg); err != nil {
		return err
	}

	path, _ := ConfigPath()
	// reuse your existing Describe styling
	// just a simple confirmation for now
	println("\nConfig saved to: " + path + "\n")

	return nil
}
