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

Optionally narrow the search by providing a subject type, a search term, or both
before launching the interactive finder. The finder includes a live preview panel
for the selected subject.

Flags:
  -d, --dest   Project directory to operate on

Examples:
  operatree find                        # browse all subjects interactively
  operatree find event                  # filter to events, then pick one
  operatree find event cairo            # filter to events matching "cairo"
  operatree find cairo                  # search "cairo" across all subject types`,
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
		fmt.Printf("Find type: %s\n", cliType)
		fmt.Printf("Find term: %s\n", cliTerm)
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

		s, err := project.FindSubjects(&p, t, term)
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
