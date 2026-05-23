package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/hanymamdouh82/operatree/internal/subject"
	"github.com/spf13/cobra"
)

var subjectName string
var subjectDate string

func init() {
	newCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	newCmd.Flags().StringVar(&subjectName, "name", "", "subject name")
	newCmd.Flags().StringVar(&subjectDate, "date", "", "subject date")
	newCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:       "new [event | task | topic | objective]",
	Short:     "Creates new subject",
	Long:      "Creates new subject within project",
	ValidArgs: SubjectValidArgs,
	Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
	Run:       newSubject,
}

var (
	argToSubject map[string]subject.SubjectType = map[string]subject.SubjectType{
		"event":     subject.SubjectEvent,
		"task":      subject.SubjectTask,
		"topic":     subject.SubjectTopic,
		"objective": subject.SubjectObjective,
	}
)

func newSubject(cmd *cobra.Command, args []string) {
	a := args[0]
	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	st, ok := argToSubject[a]
	if !ok {
		log.Fatal("unsupported subject type")
	}

	if err := project.NewSubject(&p, subjectName, subjectDate, st); err != nil {
		log.Fatal(err)
	}
}
