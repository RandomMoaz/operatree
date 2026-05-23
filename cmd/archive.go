package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	archiveCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	archiveCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(archiveCmd)
}

var archiveCmd = &cobra.Command{
	Use:   "archive [type] [term]",
	Short: "Finds a subject in project",
	Long:  "Fuzzy-Find a subject in a project",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run:   archiveSubject,
}

func archiveSubject(cmd *cobra.Command, args []string) {
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

	if err := p.Archive(s); err != nil {
		log.Fatal(err)
	}
}
