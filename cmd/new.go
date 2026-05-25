package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/hanymamdouh82/operatree/internal/subject"
	"github.com/spf13/cobra"
)

var subjectName string
var subjectDate string
var validSubjects []cobra.Completion
var newCmd = &cobra.Command{}

func init() {
	// build completion slice from available subjects dynamically
	for k := range project.SubjectModuleMap {
		sn := strings.ToLower(string(k))
		validSubjects = append(validSubjects, sn)
	}

	// define command
	newCmd = &cobra.Command{
		Use:       fmt.Sprintf("new [%s]", strings.Join(validSubjects, " | ")),
		Short:     "Creates new subject",
		Long:      "Creates new subject within project",
		ValidArgs: validSubjects,
		Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
		Run:       newSubject,
	}

	newCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	newCmd.Flags().StringVar(&subjectName, "name", "", "subject name")
	newCmd.Flags().StringVar(&subjectDate, "date", "", "subject date")
	newCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(newCmd)
}

func newSubject(cmd *cobra.Command, args []string) {
	a := args[0]

	// safety check -> dynamic loading prevents reach this error
	if a == "" {
		log.Fatal("unsupprted subject")
	}

	st := strings.ToUpper(a)

	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := project.NewSubject(&p, subjectName, subjectDate, subject.SubjectType(st)); err != nil {
		log.Fatal(err)
	}
}
