package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

func init() {
	editCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	editCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit [type] [term]",
	Short: "Edit subject metadata",
	Long: `Fuzzy-find a subject and open its META.yaml in your configured editor.

Optionally narrow the search by providing a subject type, a search term, or both
before launching the interactive finder. The project metadata index is updated
automatically once the editor is closed.

The editor is set during 'operatree init' and stored in config.
Falls back to the $EDITOR environment variable if not set in config.

Flags:
  -d, --dest   Project directory to operate on

Examples:
  operatree edit                         # browse all subjects interactively
  operatree edit task                    # filter to tasks, then pick one
  operatree edit task report             # filter to tasks matching "report"
  operatree edit -d /path/to/project`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run:  editMetadata,
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

	s, err := project.FindSubject(&p, t, term)
	if err != nil {
		log.Fatal(err)
	}

	// call edit
	if err := s.EditMetadata(); err != nil {
		log.Fatal(err)
	}

	// Sync project metadata with subject metadata
	if err := project.Sync(&p); err != nil {
		log.Fatal(err)
	}
}
