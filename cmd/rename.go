package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

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
	Long: `Fuzzy-find a subject and rename it interactively, or target one directly by UUID for scripting.

Updates the subject directory name, META.yaml, and all cross-references
in the project metadata index in one operation.

Interactive mode — launches the finder to select a subject:
  operatree rename                       # browse all subjects
  operatree rename task                  # filter to tasks, then pick one
  operatree rename task report           # filter to tasks matching "report"

Non-interactive mode — target a subject directly by UUID:
  operatree rename --uuid [uuid] --new-name [name]

Note: --uuid requires --new-name. Providing --uuid without --new-name is an error.

Flags:
  -u, --uuid       Subject UUID for non-interactive mode
  -n, --new-name   New name to assign to the subject (required with --uuid)
  -d, --dest       Project directory to operate on

Examples:
  operatree rename
  operatree rename task report
  operatree rename --uuid a1b2c3d4 --new-name "Cairo Factory Review"
  operatree rename --uuid a1b2c3d4 --new-name "Cairo Factory Review" -d /path/to/project`,
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
