package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	metadataCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	metadataCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(metadataCmd)
}

var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Edits subject metadata",
	Long:  "Opens editor to edit metadata",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run:   editMetadata,
}

func editMetadata(cmd *cobra.Command, args []string) {
	var t, term string

	if len(args) == 2 {
		t = args[0]
		term = args[1]
	} else if len(args) == 1 {
		term = args[0]
	}

	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := project.EditMetadata(&p, t, term); err != nil {
		log.Fatal(err)
	}

	if err := project.Sync(&p); err != nil {
		log.Fatal(err)
	}
}
