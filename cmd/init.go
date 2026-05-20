package cmd

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize OperaTree configuration",
	Long:  "Creates initial OperaTree config file in OS config directory",
	Args:  cobra.NoArgs,
	Run:   initConfig,
}

func initConfig(cmd *cobra.Command, args []string) {

	// Check if config already exists
	existing, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	var standardDir string
	var overwrite bool

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
			return
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
		log.Fatal(err)
	}

	cfg := config.Config{
		StandardDir: standardDir,
		Projects:    existing.Projects, // preserve tracked projects if overwriting
		Daemon: config.Daemon{
			Enabled:  false,
			Host:     "localhost",
			Port:     7070,
			DBDriver: "sqlite",
		},
	}

	if err := config.Save(cfg); err != nil {
		log.Fatal(err)
	}

	path, _ := config.ConfigPath()
	// reuse your existing Describe styling
	// just a simple confirmation for now
	println("\nConfig saved to: " + path + "\n")
}
