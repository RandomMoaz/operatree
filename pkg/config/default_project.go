package config

import (
	"fmt"
	"log"
	"slices"

	"github.com/charmbracelet/huh"
)

// prints current default project
func ShowDefulatProject() {

	cfg, err := Load()
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Default.AbsPath == "" {
		fmt.Println("No default project set. Run 'operatree default' to set one.")
		return
	}
	fmt.Printf("Default project: %s (%s)\n", cfg.Default.Name, cfg.Default.AbsPath)
}

// Uses flag -d to set default project
func SetDefaultProjectCLI(ppth string) error {

	cfg, err := Load()
	if err != nil {
		return err
	}

	if len(cfg.Projects) == 0 {
		return fmt.Errorf("No tracked projects found. Run 'operatree bootstrap' to create one.")
	}

	pidx := slices.IndexFunc(cfg.Projects, func(p Project) bool {
		return p.AbsPath == ppth
	})

	if pidx == -1 {
		return fmt.Errorf("project doesn't exist. Track project before set as default")
	}

	p := cfg.Projects[pidx]
	if err := SetDefaultProject(p); err != nil {

		return err
	}

	fmt.Printf("\nDefault project set: %s (%s)\n", p.Name, p.AbsPath)

	return nil
}

// Uses interactive CLI to set default project
func SetDefaultProjectInteractive() error {
	cfg, err := Load()
	if err != nil {
		return err
	}

	// interactive select
	if len(cfg.Projects) == 0 {
		return fmt.Errorf("No tracked projects found. Run 'operatree bootstrap' to create one.")
	}

	options := make([]huh.Option[string], len(cfg.Projects))
	for i, p := range cfg.Projects {
		label := fmt.Sprintf("%-20s  %s", p.Name, p.AbsPath)
		options[i] = huh.NewOption(label, p.AbsPath)
	}

	var selected string
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select default project").
				Description("Used automatically when -d flag is not provided").
				Options(options...).
				Value(&selected),
		),
	).Run()
	if err != nil {
		return err
	}

	for _, p := range cfg.Projects {
		if p.AbsPath == selected {
			if err := SetDefaultProject(p); err != nil {
				return err
			}
			fmt.Printf("\nDefault project set: %s (%s)\n", p.Name, p.AbsPath)
		}
	}

	return nil
}
