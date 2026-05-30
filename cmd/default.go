package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	setDPCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	setDPCmd.PreRun = resolveProjectDirSkippingConfig
	rootCmd.AddCommand(setDPCmd)
}

var setDPCmd = &cobra.Command{
	Use:   "default",
	Short: "Sets default project",
	Long:  "Sets a default project from tracked projects",
	Args:  cobra.NoArgs,
	Run:   setDefaultProject,
}

func setDefaultProject(cmd *cobra.Command, args []string) {

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
