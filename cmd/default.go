package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/spf13/cobra"
)

var showDefault bool

func init() {
	setDPCmd.Flags().BoolVar(&showDefault, "show", false, "show current default project")
	setDPCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	setDPCmd.PreRun = resolveProjectDirSkippingConfig
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

	if showDefault {
		config.ShowDefulatProject()
		return
	}

	if destDir != "" {
		if err := config.SetDefaultProjectCLI(actDir); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := config.SetDefaultProjectInteractive(); err != nil {
		log.Fatal(err)
	}
}
