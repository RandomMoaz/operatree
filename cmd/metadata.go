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
	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	var t, term string

	if len(args) == 2 {
		t = args[0]
		term = args[1]
	} else if len(args) == 1 {
		term = args[0]
	} else {
		t = ""
		term = ""
	}

	s, err := project.FindSubjects(&p, t, term)
	if err != nil {
		log.Fatal(err)
	}

	// call edit
	if err := project.EditSubjectMetadata(&p, s); err != nil {
		log.Fatal(err)
	}

	// Sync project metadata with subject metadata
	if err := project.Sync(&p); err != nil {
		log.Fatal(err)
	}
}
