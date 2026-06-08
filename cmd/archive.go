package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

func init() {
	archiveCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	archiveCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(archiveCmd)
}

var archiveCmd = &cobra.Command{
	Use:   "archive [type] [term]",
	Short: "Archive a subject",
	Long: `Fuzzy-find a subject and move it to the project archive (99_ARCHIVE/).

Optionally narrow the search by providing a subject type and/or search term
before launching the interactive finder.

Examples:
  operatree archive                   # browse all subjects interactively
  operatree archive task              # filter to tasks, then pick one
  operatree archive task report       # filter to tasks matching "report"`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run:  archiveSubject,
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

	s, err := project.FindSubject(&p, t, term)
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Archive(s); err != nil {
		log.Fatal(err)
	}
}
