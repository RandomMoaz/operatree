package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	syncCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	syncCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Syncs project",
	Long:  "Syncs project subjects with project metadata",
	Args:  cobra.NoArgs,
	Run:   sync,
}

func sync(cmd *cobra.Command, args []string) {
	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	// Sync project metadata with subject metadata
	if err := project.Sync(&p); err != nil {
		log.Fatal(err)
	}
}
