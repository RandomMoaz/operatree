package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/hanymamdouh82/operatree/internal/runner"
	"github.com/spf13/cobra"
)

func init() {
	openSubjectCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	openSubjectCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(openSubjectCmd)
}

var openSubjectCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens subject directory",
	Long:  "Opens subject directory in default file manager",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run:   openSubject,
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

	s, err := project.FindSubjects(&p, t, term)
	if err != nil {
		log.Fatal(err)
	}

	if err := runner.OpenFileManager(s.DirName); err != nil {
		log.Fatal(err)
	}
}
