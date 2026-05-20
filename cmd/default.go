package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/spf13/cobra"
)

var showDefault bool

func init() {
	setDPCmd.Flags().BoolVar(&showDefault, "show", false, "show current default project")
	rootCmd.AddCommand(setDPCmd)
}

var setDPCmd = &cobra.Command{
	Use:   "default",
	Short: "Set or show default project",
	Long:  "Sets a default project from tracked projects, or shows the current default",
	Args:  cobra.NoArgs,
	Run:   setDefaultProject,
}

func setDefaultProject(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// --show flag
	if showDefault {
		if cfg.Default.AbsPath == "" {
			fmt.Println("No default project set. Run 'operatree default' to set one.")
			return
		}
		fmt.Printf("Default project: %s (%s)\n", cfg.Default.Name, cfg.Default.AbsPath)
		return
	}

	// interactive select
	if len(cfg.Projects) == 0 {
		fmt.Println("No tracked projects found. Run 'operatree bootstrap' to create one.")
		return
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
		log.Fatal(err)
	}

	for _, p := range cfg.Projects {
		if p.AbsPath == selected {
			if err := config.SetDefaultProject(p); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\nDefault project set: %s (%s)\n", p.Name, p.AbsPath)
			return
		}
	}
}
