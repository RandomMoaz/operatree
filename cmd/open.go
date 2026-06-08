package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/runner"
	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

func init() {
	openSubjectCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	openSubjectCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(openSubjectCmd)
}

var openSubjectCmd = &cobra.Command{
	Use:   "open [type] [term]",
	Short: "Open a subject directory",
	Long: `Fuzzy-find a subject and open its directory in your configured file manager.

Optionally narrow the search by providing a subject type, a search term, or both
before launching the interactive finder.

The file manager is set during 'operatree init' and stored in config.

Examples:
  operatree open                         # browse all subjects interactively
  operatree open task                    # filter to tasks, then pick one
  operatree open task report             # filter to tasks matching "report"
  operatree open -d /path/to/project     # browse all subjects interactively in project at /path/to/project`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run:  openSubject,
}

func openSubject(cmd *cobra.Command, args []string) {
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

	if err := runner.OpenFileManager(s.DirName); err != nil {
		log.Fatal(err)
	}
}
