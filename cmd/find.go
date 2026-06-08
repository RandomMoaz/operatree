package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/hanymamdouh82/operatree/pkg/subject"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var cliTerm, cliType string
var isPlain bool

func init() {
	findCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	findCmd.Flags().StringVarP(&cliTerm, "term", "t", "", "term")
	findCmd.Flags().StringVarP(&cliType, "type", "s", "", "subject type")
	findCmd.Flags().BoolVarP(&isPlain, "plain", "p", false, "show plain result")

	findCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find [type] [term]",
	Short: "Find a subject",
	Long: `Fuzzy-find subjects across all metadata fields — name, tags, participants, notes, date, and location.

Use positional arguments for interactive mode, or --term and --type flags for
non-interactive scripting. The --plain flag outputs raw YAML in both modes.

Interactive mode — launches the finder with a live preview panel:
  operatree find                        # browse all subjects
  operatree find event                  # filter to events, then pick one
  operatree find event cairo            # filter to events matching "cairo"

Non-interactive mode — returns a list of matching subjects directly:
  operatree find --term cairo           # search all subject types for "cairo"
  operatree find --term cairo --type event   # search events only

Flags:
  -t, --term    Search term for non-interactive mode
  -s, --type    Subject type filter for non-interactive mode
  -p, --plain   Output results as raw YAML instead of formatted view
  -d, --dest    Project directory to operate on

Examples:
  operatree find
  operatree find event cairo
  operatree find --term cairo --plain
  operatree find --term report --type task --plain | grep owner
  operatree find --term cairo -d /path/to/project`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run:  find,
}

func find(cmd *cobra.Command, args []string) {

	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	// non-interactive prompt
	if cliTerm != "" {
		ss, err := project.FindSubjectsSilent(&p, cliType, cliTerm)
		if err != nil {
			log.Fatal(err)
		}

		for _, s := range ss {
			if err := showResult(s, isPlain); err != nil {
				log.Fatal(err)
			}
		}

	} else {
		// interactive prompt
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

		if s.Type != "" {
			if err := showResult(s, isPlain); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func showResult(s subject.Subject, isPlain bool) error {

	if !isPlain {
		s.Describe()
		return nil
	}

	b, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", b)
	return nil
}
