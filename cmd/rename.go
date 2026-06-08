package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

var newName, uuid string

func init() {
	renameSubjectCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	renameSubjectCmd.Flags().StringVarP(&newName, "new-name", "n", "", "subject new name")
	renameSubjectCmd.Flags().StringVarP(&uuid, "uuid", "u", "", "subject UUID")
	renameSubjectCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(renameSubjectCmd)
}

var renameSubjectCmd = &cobra.Command{
	Use:   "rename [type] [term]",
	Short: "Rename a subject",
	Long: `Fuzzy-find a subject and rename it interactively.

Updates both the subject directory name and its META.yaml in place.
Optionally narrow the search by providing a subject type, a search term, or both
before launching the interactive finder.

Flags:
  -d, --dest   Project directory to operate on

Examples:
  operatree rename                       # browse all subjects interactively
  operatree rename task                  # filter to tasks, then pick one
  operatree rename task report           # filter to tasks matching "report"
  operatree rename -d /path/to/project`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run:  renameSubject,
}

func renameSubject(cmd *cobra.Command, args []string) {

	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	// uuid cannot work alone, new name must be provided
	if uuid != "" && newName == "" {
		log.Fatal(fmt.Errorf("cannot rename without new-name flag"))
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

	if err := p.RenameSubject(t, term, newName, uuid); err != nil {
		log.Fatal(err)
	}
}
