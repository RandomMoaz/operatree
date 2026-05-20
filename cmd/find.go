package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Finds a subject in project",
	Long:  "Fuzzy-Find a subject in a project",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(2), cobra.MinimumNArgs(1)),
	Run:   find,
}

func find(cmd *cobra.Command, args []string) {

	p, err := project.Load(prjDir)
	if err != nil {
		log.Fatal(err)
	}

	var t, term string

	if len(args) == 2 {
		t = args[0]
		term = args[1]
	} else {
		term = args[0]
	}

	s, err := project.FindSubjects(&p, t, term)
	if err != nil {
		log.Fatal(err)
	}

	if s.Type != "" {
		s.Describe()
	}
}
